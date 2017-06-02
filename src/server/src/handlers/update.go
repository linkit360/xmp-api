package handlers

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
)

func Update() {
	log.Info("Handlers: Update")

	for {
		log.Info("Handlers: Update: z1")
		update := <-base.ChanUpdate
		log.Info("Handlers: Update: ", update.For)

		jsval, err := json.Marshal(update)
		if err != nil {
			log.Error("Handlers: Update: ", err)
		}
		log.Info("Handlers: Update: z2")

		Send(update.For, jsval)
		log.Info("Handlers: Update: z3")
	}
}
