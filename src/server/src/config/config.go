package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

type DbConfig struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type AwsConfig struct {
	Id     string `json:"key"`
	Secret string `json:"secret"`
}

func LoadConfig() (DbConfig, AwsConfig) {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)
	var err error

	var envStr = "/config/" + EnvString("PROJECT_ENV", "development")

	// DB Config
	dat, err := ioutil.ReadFile(envStr + "/db.json")
	if err != nil {
		log.WithFields(
			log.Fields{
				"prefix": "Config",
				"error":  err.Error(),
			},
		).Fatal("Read Error")
	}

	var cfg DbConfig
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		log.WithFields(
			log.Fields{
				"prefix": "Config",
				"error":  err.Error(),
			},
		).Fatal("Load Error")
	}

	// AWS Config
	dat, err = ioutil.ReadFile(envStr + "/aws.json")
	if err != nil {
		log.WithFields(
			log.Fields{
				"prefix": "Config",
				"error":  err.Error(),
			},
		).Fatal("Read Error")
	}

	var aws AwsConfig
	err = json.Unmarshal(dat, &aws)
	if err != nil {
		log.WithFields(
			log.Fields{
				"prefix": "Config",
				"error":  err.Error(),
			},
		).Fatal("Load Error")
	}

	log.WithFields(log.Fields{
		"prefix": "Config",
		"env":    EnvString("PROJECT_ENV", "development"),
	}).Info("Init Done")

	return cfg, aws
}

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
