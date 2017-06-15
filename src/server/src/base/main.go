package base

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/jinzhu/gorm.v1"

	"github.com/linkit360/xmp-api/src/server/src/config"
)

var cfg config.DbConfig
var cfgAws config.AwsConfig
var db *gorm.DB
var ChanUpdate chan UpdateCall

func init() {
	ChanUpdate = make(chan UpdateCall)
}

func Init() {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)
	var err error
	cfg, cfgAws = config.LoadConfig()

	connstr := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host,
		cfg.User,
		cfg.Database,
		cfg.Password,
	)

	db, err = gorm.Open(
		"postgres",
		connstr,
	)

	if err != nil {
		log.Panic("Base: Connection Error ", err.Error())
	}
}
