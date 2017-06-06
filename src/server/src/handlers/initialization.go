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
			log.Info("Initialization: " + instance_id + ": " + Clients[instance_id])

			out.Ok = true
			out.Error = ""

			// Load Services for instance
			out.Services, err = base.GetServices(id_provider)
			if err != nil {
				out.Error = err.Error()
			}

			if len(out.Services) > 0 {
				// Load Campaigns for instance
				out.Campaigns, err = base.GetCampaigns(out.Services)
				if err != nil {
					out.Error = err.Error()
				}
			}
		} else {
			out.Error = "Status not 1"
		}
	} else {
		out.Error = "Provider not found"
	}

	if out.Error != "" {
		out.Campaigns = nil
		out.Services = nil

		out.Ok = false
		out.Error = "Init: " + out.Error
	}

	c.JSON(
		200,
		out,
	)
}
