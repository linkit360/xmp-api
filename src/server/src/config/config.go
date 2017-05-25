package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

type DbConfig struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfig() DbConfig {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)
	var err error

	var envConfigFile string = "/config/" + EnvString("PROJECT_ENV", "development") + "/db.json"
	dat, err := ioutil.ReadFile(envConfigFile)
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

	log.WithFields(log.Fields{
		"prefix": "Config",
		"env":    EnvString("PROJECT_ENV", "development"),
	}).Info("Init Done")

	return cfg
}

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
