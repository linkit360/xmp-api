package handlers

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
)

func Update() {
	for {
		update := <-base.ChanUpdate
		log.Info("Handlers: Update: ", update.For)

		jsval, err := json.Marshal(update)
		if err != nil {
			log.Error("Handlers: Update: ", err)
		}

		Send(update.For, jsval)
	}
}
