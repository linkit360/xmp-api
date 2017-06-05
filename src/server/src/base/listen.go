package base

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/lib/pq"
)

type UpdateCall struct {
	Type string `json:"type"`
	For  string `json:"for"`
	Data string `json:"data,omitempty"`
}

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

	for {
		waitForNotification(listener)
	}
}

func waitForNotification(l *pq.Listener) {
	select {
	case payload := <-l.Notify:
		handleNotify(payload.Extra)

	case <-time.After(90 * time.Second):
		go l.Ping()
	}
}

func handleNotify(payload string) {
	pl := UpdateCall{}
	err := json.Unmarshal([]byte(payload), &pl)
	if err != nil {
		log.Error("Listen: waitForNotification: ", err)
		return
	}

	log.Info("Listen: " + pl.Type + " for " + pl.For)
	ChanUpdate <- pl
	log.Info("Listen: Write OK")
}
