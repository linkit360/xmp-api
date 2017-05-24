package xmp_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

var client *http.Client
var config ClientConfig

type ClientConfig struct {
	Enabled    bool   `yaml:"enabled"`
	DSN        string `default:":50307" yaml:"dsn"`
	Timeout    int    `default:"10" yaml:"timeout"`
	InstanceId string `default:"" yaml:"instance_id"`
}

func Init(clientConf ClientConfig) error {
	log.SetLevel(log.DebugLevel)

	config = clientConf
	client = &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}

	//runGin()
	return nil
}

/*
func runGin() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET(
		"/ping",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		},
	)

	// use config
	r.Run(":50318")
}
*/

func Call(funcName string, res interface{}, req ...interface{}) error {
	if !config.Enabled {
		return fmt.Errorf("Acceptor Client Disabled")
	}

	var url string = "http://" + config.DSN + "/" + funcName + "?instance_id=" + config.InstanceId
	var err error

	// GET by default
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if len(req) > 0 {
		// POST
		jsonValue, err := json.Marshal(req)
		if err != nil {
			return err
		}

		request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		if err != nil {
			return err
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	return nil
}
