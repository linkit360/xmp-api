package xmp_api_client

import (
	"encoding/json"
	"math/rand"
	"time"

	log "github.com/Sirupsen/logrus"
	xmp_api_structs "github.com/linkit360/xmp-api/src/structs"
	"gopkg.in/gin-gonic/gin.v1"
)

type UpdateRequest struct {
	Type string `json:"type"` // Type of entity (service/campaign) && type of event (new/update) ex: service.new
	For  string `json:"for"`  // Instance ID who is responsible for handling the event
	Data string `json:"data"` // JSON of entity
}

func GetRandomAggregate() xmp_api_structs.Aggregate {
	return xmp_api_structs.Aggregate{
		InstanceId:             "a7da1e9f-fcc1-4087-9c58-4d31bcdbd515",
		ReportAt:               time.Now().UTC().Unix(),
		CampaignCode:           "290",
		OperatorCode:           52000,
		LpHits:                 rand.Int63n(200),
		LpMsisdnHits:           rand.Int63n(150),
		MoTotal:                rand.Int63n(200),
		MoChargeSuccess:        rand.Int63n(200),
		MoChargeSum:            1000.,
		MoChargeFailed:         rand.Int63n(200),
		MoRejected:             rand.Int63n(200),
		RenewalTotal:           rand.Int63n(200),
		RenewalChargeSuccess:   rand.Int63n(200),
		RenewalChargeSum:       12312.,
		RenewalFailed:          rand.Int63n(200),
		InjectionTotal:         rand.Int63n(200),
		InjectionChargeSuccess: rand.Int63n(200),
		Pixels:                 rand.Int63n(200),
	}
}

func update(c *gin.Context) {
	req := UpdateRequest{}
	c.BindJSON(&req)

	if req.Type == "service.new" || req.Type == "service.update" {
		svc := xmp_api_structs.Service{}
		err := json.Unmarshal([]byte(req.Data), &svc)
		if err != nil {
			log.WithFields(log.Fields{
				"prefix": "XMPAPI",
			}).Error(
				"Update: ",
				err.Error(),
			)
		}

		log.WithFields(log.Fields{
			"prefix": "XMPAPI",
			"type":   req.Type,
		}).Info(svc.Id)

		ChanServices <- svc
	}

	if req.Type == "campaign.new" || req.Type == "campaign.update" {
		campaign := xmp_api_structs.Campaign{}
		err := json.Unmarshal([]byte(req.Data), &campaign)
		if err != nil {
			log.Error("Update: ", err)
		}

		log.WithFields(log.Fields{
			"prefix": "XMPAPI",
			"type":   req.Type,
		}).Info(campaign.Id)

		ChanCampaigns <- campaign
	}

	log.WithFields(log.Fields{
		"prefix": "XMPAPI",
		"type":   req.Type,
	}).Info("Update OK")

	c.JSON(
		200,
		gin.H{
			"message": "ok",
		},
	)
}
