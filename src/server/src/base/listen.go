package base

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/lib/pq"
)

var ChanUpdate chan updateCall

func Listen() {
	log.Info("Base: Listen")
	var err error

	connstr := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host,
		cfg.User,
		cfg.Database,
		cfg.Password,
	)

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	listener := pq.NewListener(connstr, 10*time.Second, time.Minute, reportProblem)
	err = listener.Listen("xmp_update")
	if err != nil {
		panic(err)
	}
	fmt.Println("entering main loop")
	for {
		waitForNotification(listener)
	}
}

func waitForNotification(l *pq.Listener) {
	select {
	case payload := <-l.Notify:
		fmt.Println("received notification, new work available: ", payload.Extra)

		pl := updateCall{}
		err := json.Unmarshal([]byte(payload.Extra), &pl)
		if err != nil {
			log.Error("Listen: waitForNotification: ", err)
		} else {
			log.Error("Listen: " + pl.Type + " for " + pl.For)

			pl.Payload = payload.Extra
			ChanUpdate <- pl
		}

	case <-time.After(90 * time.Second):
		go l.Ping()
		// Check if there's more work available, just in case it takes
		// a while for the Listener to notice connection loss and
		// reconnect.
		fmt.Println("received no work for 90 seconds, checking for new work")
	}
}

type updateCall struct {
	Type    string `json:"type"`
	For     string `json:"for"`
	Payload string `json:"payload,omitempty"`
}
