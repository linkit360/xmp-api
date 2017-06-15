package handlers

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/structs"
)

func Initialization(c *gin.Context) {
	var err error
	var instance_id string = c.Query("instance_id")
	var out xmp_api_structs.HandShake
	out.Ok = true
	out.Error = ""

	status, id_provider := base.GetOptions(instance_id)
	if id_provider > 0 {
		// Found instance
		if status == 1 {
			// save client
			Clients[instance_id] = c.ClientIP()

			// Load Services for instance
			out.Services, err = base.GetServices(id_provider)
			if err != nil {
				out.Error = "Services: " + err.Error()
			}

			if len(out.Services) > 0 {
				// Load Campaigns for instance
				out.Campaigns, err = base.GetCampaigns(out.Services)
				if err != nil {
					out.Error = "Campaigns: " + err.Error()
				}
			} else {
				out.Error = "No Services"
			}

			// Load Country for instance
			out.Country, err = base.GetCountry(id_provider)
			if err != nil {
				out.Error = "Country: " + err.Error()
			}

			// Load Blacklist for instance
			out.BlackList, err = base.GetBlacklist(id_provider)
			if err != nil {
				out.Error = "BlackList: " + err.Error()
			}

			// Load Blacklist for instance
			out.Operators, err = base.GetOperators(id_provider)
			if err != nil {
				out.Error = "Operators: " + err.Error()
			}
		} else {
			out.Error = "Instance status " + strconv.Itoa(status)
		}
	} else {
		out.Error = "Provider not found"
	}

	var logmsg string = "INIT | IID: " + instance_id +
		" | PROV: " + strconv.Itoa(id_provider) +
		" | SVCs: " + strconv.Itoa(len(out.Services)) +
		" | CAMPs: " + strconv.Itoa(len(out.Campaigns)) +
		" | IP: " + Clients[instance_id]

	if out.Error != "" {
		logmsg = logmsg + " | Error: " + out.Error

		out.Campaigns = nil
		out.Services = nil

		out.Ok = false
		out.Error = "Init: " + out.Error
	}

	log.Info(logmsg)

	c.JSON(
		200,
		out,
	)
}
