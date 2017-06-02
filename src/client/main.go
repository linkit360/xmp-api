package xmp_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/structs"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/gin-gonic/gin.v1"
)

var client *http.Client
var config ClientConfig
var ChanServices chan xmp_api_structs.Service

type ClientConfig struct {
	Enabled    bool   `yaml:"enabled"`
	DSN        string `default:":50307" yaml:"dsn"`
	Timeout    int    `default:"10" yaml:"timeout"`
	InstanceId string `default:"" yaml:"instance_id"`
}

func init() {
	ChanServices = make(chan xmp_api_structs.Service, 1)
}

func Init(clientConf ClientConfig) error {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)

	config = clientConf
	client = &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}

	log.Info("XMPAPI: Init")
	go runGin()
	return nil
}

func runGin() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/update", update)

	// use config
	r.Run(":50319")
}

func Call(funcName string, res interface{}, req ...interface{}) error {
	if !config.Enabled {
		return fmt.Errorf("Acceptor Client Disabled")
	}

	var url string = "http://" + config.DSN + "/" + funcName + "?instance_id=" + config.InstanceId
	var err error

	// GET by default
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %s", err.Error())
	}

	if len(req) > 0 {
		// POST
		jsonValue, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("json.Marshal: %s", err.Error())
		}

		//log.Debug(string(jsonValue))

		request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		if err != nil {
			return fmt.Errorf("POST http.NewRequest: %s", err.Error())
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("client.Do: %s", err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return fmt.Errorf("%s, json.Unmarshal: %s", string(body), err.Error())
	}

	return nil
}
