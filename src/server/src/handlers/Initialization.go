package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/structs"
	"gopkg.in/gin-gonic/gin.v1"
)

func Initialization(c *gin.Context) {
	log.Info("")
	log.Info("Call Initialization")
	var instance_id string = c.Query("instance_id")
	log.Info(instance_id)

	var err error

	status, id_provider := base.GetOptions(instance_id)
	out := xmp_api_structs.HandShake{
		Ok:     false,
		Error:  "Status not 1",
		Status: status,
	}

	if out.Status == 1 {
		// save client
		Clients[instance_id] = c.ClientIP()

		out.Ok = true
		out.Error = ""
		out.ProviderId = id_provider
		out.Services, err = base.GetServices(id_provider)
		if err != nil {
			out.Ok = false
			out.Error = err.Error()
			out.Services = nil
		}
	}

	// remove me
	//Send("2f4fd741-61ef-45ab-8436-840ce54d6d29")

	c.JSON(
		200,
		out,
	)
}
