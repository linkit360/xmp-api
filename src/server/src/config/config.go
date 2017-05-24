package config

import (
	"flag"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/configor"
	"github.com/linkit360/go-utils/db"
)

type ServerConfig struct {
	Port string `default:"40400" yaml:"port"`
}

type AppConfig struct {
	AppName string            `yaml:"app_name"`
	Server  ServerConfig      `yaml:"server"`
	DbConf  db.DataBaseConfig `yaml:"db"`
}

func LoadConfig() AppConfig {
	var envConfigFile string = "/config/acceptor." + EnvString("PROJECT_ENV", "development") + ".yml"

	cfg := flag.String("config", envConfigFile, "configuration yml file")
	flag.Parse()
	var appConfig AppConfig
	if *cfg != "" {
		if err := configor.Load(&appConfig, *cfg); err != nil {
			log.WithFields(log.Fields{
				"prefix": "Config",
				"error":  err.Error(),
			}).Fatal("Load Error")
			os.Exit(1)
		}
	}

	if appConfig.AppName == "" {
		log.Fatal("app name must be defiled as <host>-<name>")
	}

	if strings.Contains(appConfig.AppName, "-") {
		log.Fatal("app name must be without '-' : it's not a valid metric name")
	}

	log.WithFields(log.Fields{
		"prefix": "Config",
		"ENV":    EnvString("PROJECT_ENV", "development"),
	}).Info("Init Done")

	return appConfig
}

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
