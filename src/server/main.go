package xmp_api_server

import (
	"runtime"

	logr "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/server/src/handlers"
	"github.com/linkit360/xmp-api/src/server/src/websocket"
)

func Init() {
	logr.SetFormatter(new(prefixed.TextFormatter))
	logr.SetLevel(logr.DebugLevel)
	log := logr.WithFields(logr.Fields{
		"prefix": "Main",
	})

	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	//log.WithField("CPUCount", nuCPU)

	base.Init()
	go websocket.Init()

	go base.Listen()
	go handlers.Update()

	log.Info("Init Done")
	runGin()
}

func runGin() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/initialization", handlers.Initialization)
	r.POST("/aggregate", handlers.Aggregate)

	r.Run(":50318")
}
