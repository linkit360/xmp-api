package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/structs"
	"gopkg.in/gin-gonic/gin.v1"
)

func Initialization(c *gin.Context) {
	var err error
	var instance_id string = c.Query("instance_id")
	var out xmp_api_structs.HandShake

	log.Info("Call Initialization: " + instance_id)

	status, id_provider := base.GetOptions(instance_id)
	if id_provider > 0 {
		// Found instance
		if status == 1 {
			// save client
			Clients[instance_id] = c.ClientIP()

			out.Ok = true
			out.Error = ""

			// Load Services for instance
			out.Services, err = base.GetServices(id_provider)
			if err != nil {
				out.Ok = false
				out.Error = err.Error()
				out.Services = nil
			}
		} else {
			out.Ok = false
			out.Error = "Status not 1"
		}
	} else {
		out.Ok = false
		out.Error = "Provider not found"
	}

	if out.Error != "" {
		out.Error = "Init: " + out.Error
	}

	// remove me
	//Send("2f4fd741-61ef-45ab-8436-840ce54d6d29")

	c.JSON(
		200,
		out,
	)
}
