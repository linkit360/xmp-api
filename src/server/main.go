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
	//log.WithField("port", appConfig.Server.HttpPort).Info("service port")
}

/*
func runRPC(appConfig config.AppConfig) {
	l, err := net.Listen("tcp", "0.0.0.0:"+appConfig.Server.RPCPort)
	if err != nil {
		log.WithFields(log.Fields{
			"prefix": "RPC",
		}).Fatal("netListen ", err.Error())
	}

	log.WithFields(log.Fields{
		"prefix": "RPC",
		"Port":   appConfig.Server.RPCPort,
	}).Info()

	server := rpc.NewServer()
	server.RegisterName("Aggregate", &handlers.Aggregate{})
	server.RegisterName("BlackList", &handlers.BlackList{})
	server.RegisterName("Campaigns", &handlers.Campaigns{})

	for {
		if conn, err := l.Accept(); err == nil {

			log.WithFields(log.Fields{
				"prefix": "RPC",
				"local":  conn.LocalAddr().String(),
				"remote": conn.RemoteAddr().String(),
			}).Info("CONNECT!")

			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		} else {
			log.WithFields(log.Fields{
				"prefix": "RPC",
			}).Info("accept", err.Error())
		}
	}
}
*/
