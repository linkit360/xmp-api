package handlers

import (
	log "github.com/Sirupsen/logrus"

	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/server/src/websocket"
	"github.com/linkit360/xmp-api/src/structs"
	"gopkg.in/gin-gonic/gin.v1"
)

func Aggregate(c *gin.Context) {
	var err error
	var total int64
	var instance_id string = c.Query("instance_id")

	log.Infoln()
	log.Info("Call Aggregate: " + instance_id)

	items := []xmp_api_structs.Aggregate{}
	err = c.BindJSON(&items)
	if err == nil {
		//log.Debugf("%#+v\n", items)
		websocket.NewReports(items)
		err, total = base.SaveRows(items)
		if err != nil {
			log.Error("Aggregate Save: ", err)
		}
	} else {
		log.Error("Aggregate Bind: ", err)
	}

	out := gin.H{}
	if err != nil {
		out["ok"] = false
		out["error"] = err.Error()
		log.Info("Aggregate FAIL")
	} else {
		out["ok"] = true
		log.Info("Aggregate OK: ", total)
	}

	c.JSON(
		200,
		out,
	)
	log.Infoln()
}
