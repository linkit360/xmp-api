package websocket

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/linkit360/xmp-api/src/server/src/base"
	"github.com/linkit360/xmp-api/src/server/src/config"
	"github.com/linkit360/xmp-api/src/structs"
	"github.com/x-cray/logrus-prefixed-formatter"
)

var lastResetTime int = 0
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Data struct {
	LpHits     uint64            `json:"lp"`
	Mo         uint64            `json:"mo"`
	MoSuccess  uint64            `json:"mos"`
	Countries  map[string]uint64 `json:"countries"`
	Logs       []string          `json:"logs"`
	ClientsCnt uint              `json:"clientsCnt"`
}

var data = Data{}
var provs = map[string]string{}

func Init() {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)

	reset()
	go resetDay()

	http.HandleFunc("/echo", echo)

	env := config.EnvString("PROJECT_ENV", "development")

	log.WithFields(log.Fields{
		"prefix": "WS",
		"env":    env,
	}).Info("Init Done")

	//if env == "development" {
	log.Fatal(http.ListenAndServe(":2082", nil))
	//} else {
	//	log.Fatal(http.ListenAndServeTLS(":2082", "/config/ssl/crt.crt", "/config/ssl/priv.key", nil))
	//}
}

func echo(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"prefix": "WS",
	}).Info("Connect")

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"prefix": "WS",
			"error":  err,
		}).Error("Upgrade")
		return
	}
	defer c.Close()

	data.ClientsCnt = data.ClientsCnt + 1

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(prepData()))
			if err != nil {
				//log.Println("WS:","write: ", err)
				c.Close()
				data.ClientsCnt = data.ClientsCnt - 1
				return
			}
		}
	}

	/*
		go func() {
			for range time.Tick(time.Second) {
				err := c.WriteMessage(websocket.TextMessage, []byte(prepData()))
				if err != nil {
					log.WithFields(log.Fields{
						"prefix": "WS",
						"error":  err,
					}).Error("Write")

					//c.Close()
					data.ClientsCnt = data.ClientsCnt - 1
				}
			}
		}()
	*/
}

func NewReports(rows []xmp_api_structs.Aggregate) {
	for _, row := range rows {
		log.WithFields(log.Fields{
			"prefix":                 "WS",
			"ProviderName":           row.ProviderName,
			"OperatorCode":           row.OperatorCode,
			"CampaignId":             row.CampaignCode,
			"LpHits":                 row.LpHits,
			"LpMsisdnHits":           row.LpMsisdnHits,
			"MoTotal":                row.MoTotal,
			"MoChargeSuccess":        row.MoChargeSuccess,
			"MoChargeSum":            row.MoChargeSum,
			"MoChargeFailed":         row.MoChargeFailed,
			"MoRejected":             row.MoRejected,
			"RenewalTotal":           row.RenewalTotal,
			"RenewalChargeSuccess":   row.RenewalChargeSuccess,
			"RenewalChargeSum":       row.RenewalChargeSum,
			"RenewalFailed":          row.RenewalFailed,
			"InjectionTotal":         row.InjectionTotal,
			"InjectionChargeSuccess": row.InjectionChargeSuccess,
			"InjectionChargeSum":     row.InjectionChargeSum,
			"InjectionFailed":        row.InjectionFailed,
			"ExpiredTotal":           row.ExpiredTotal,
			"ExpiredChargeSuccess":   row.ExpiredChargeSuccess,
			"ExpiredChargeSum":       row.ExpiredChargeSum,
			"ExpiredFailed":          row.ExpiredFailed,
			"Pixels":                 row.Pixels,
		}).Info("New Report")

		if len(data.Countries) == 0 {
			data.Countries = make(map[string]uint64, 10)
		}

		data.LpHits = data.LpHits + uint64(row.LpHits)
		data.Mo = data.Mo + uint64(row.MoTotal)
		data.MoSuccess = data.MoSuccess + uint64(row.MoChargeSuccess)
		data.Countries[provs[row.ProviderName]] =
			data.Countries[provs[row.ProviderName]] +
				uint64(row.LpHits)
	}
}

func resetDay() {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for t := range ticker.C {
			if lastResetTime != t.Day() {
				reset()
				log.WithFields(log.Fields{
					"prefix": "WS",
				}).Info("Reset Day")
			}
		}
	}()
}

func reset() {
	data.Countries, provs, data.LpHits, data.Mo, data.MoSuccess = base.GetWsData()
	lastResetTime = time.Now().Day()
}

func prepData() string {
	mapB, _ := json.Marshal(data)
	return string(mapB)
}
