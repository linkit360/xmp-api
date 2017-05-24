package xmp_api_server

import (
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/server/src/config"
	"github.com/linkit360/xmp-api/src/server/src/handlers"
	"github.com/linkit360/xmp-api/src/server/src/websocket"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/gin-gonic/gin.v1"
)

var appConfig config.AppConfig

func Init() {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)

	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	//log.WithField("CPUCount", nuCPU)

	appConfig = config.LoadConfig()

	base.Init(appConfig.DbConf)
	go websocket.Init()

	log.WithFields(log.Fields{
		"prefix": "Main",
	}).Info("Init Done")

	runGin()
}

func runGin() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/initialization", handlers.Initialization)
	r.POST("/aggregate", handlers.Aggregate)

	r.Run(":" + appConfig.Server.Port)
}
