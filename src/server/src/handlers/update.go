package handlers

import (
	"github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/server/src/base"
)

func Update() {
	logrus.Info("Handlers: Update")
	for update := range base.ChanUpdate {
		logrus.Info("Handlers: Update: ", update.For)
		Send(update.For, update.Payload)
	}
}
