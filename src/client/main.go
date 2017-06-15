package xmp_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/linkit360/xmp-api/src/structs"
)

var (
	client        *http.Client
	config        ClientConfig
	ChanServices  chan xmp_api_structs.Service
	ChanCampaigns chan xmp_api_structs.Campaign
	ChanBlacklist chan xmp_api_structs.Blacklist
)

type ClientConfig struct {
	Enabled    bool   `yaml:"enabled"`
	DSN        string `default:":50318" yaml:"dsn"`
	ClientPort int    `default:"50319" yaml:"client_port"`
	Timeout    int    `default:"10" yaml:"timeout"`
	InstanceId string `default:"" yaml:"instance_id"`
}

func init() {
	ChanServices = make(chan xmp_api_structs.Service)
	ChanCampaigns = make(chan xmp_api_structs.Campaign)
	ChanBlacklist = make(chan xmp_api_structs.Blacklist)
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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/update", update)

	// use config
	r.Run(":" + strconv.Itoa(config.ClientPort))
}

func Call(funcName string, res interface{}, req ...interface{}) error {
	if !config.Enabled {
		return fmt.Errorf("Acceptor Client Disabled")
	}

	var url string = "http://" + config.DSN + "/" + funcName + "?instance_id=" + config.InstanceId + "&port=" + strconv.Itoa(config.ClientPort)
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
