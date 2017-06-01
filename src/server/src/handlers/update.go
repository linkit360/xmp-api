package handlers

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
)

func Update() {
	log.Info("Handlers: Update")
	for update := range base.ChanUpdate {
		log.Info("Handlers: Update: ", update.For)

		jsval, err := json.Marshal(update)
		if err != nil {
			log.Error("Handlers: Update: ", err)
		}

		Send(update.For, jsval)
	}
}
