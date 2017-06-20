package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"

	"github.com/linkit360/xmp-api/src/client"
	"github.com/linkit360/xmp-api/src/server"
	"github.com/linkit360/xmp-api/src/structs"
)

func main() {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)

	//for testing client && server
	arg := os.Args[1:]

	if len(arg) < 1 {
		log.Error("No package specified.")
		return
	}

	if arg[0] == "server" {
		runServer()
	}

	if arg[0] == "client" {
		runClient()
	}
}

func runServer() {
	log.Info("Testing server")
	xmp_api_server.Init()
}

func runClient() {
	log.Info("Testing client")
	cfg := xmp_api_client.ClientConfig{
		Enabled:    true,
		DSN:        "go:50318",
		ClientPort: 50319,
		Timeout:    10,
		//InstanceId: "ecbe1211-3b1e-4c91-96fc-574dc979668a", // qrtech
		InstanceId: "967eda58-0e47-4bd4-8128-b30461eb9b19", // beeline
	}

	if err := xmp_api_client.Init(cfg); err != nil {
		log.Panic("Cannot init acceptor client", err)
	}
	log.Info("Ready")

	testInitialization()
	testAggregate()
	//testUpdateRead()
}

func testInitialization() {
	var resp xmp_api_structs.HandShake
	err := xmp_api_client.Call("initialization", &resp)
	if err != nil {
		log.Error(err)
	}

	log.WithFields(
		log.Fields{
			"prefix":    "testInitialization",
			"ok":        resp.Ok,
			"error":     resp.Error,
			"campaigns": len(resp.Campaigns),
			"services":  len(resp.Services),
			"blacklist": resp.BlackList,
		},
	).Info("Request Done")

	//log.Info("TEST! ", resp.Services)
	//log.Info("TEST! ", resp.Services["c4bc3983-648c-45f9-a5b3-e1f6aa82db90"].Price)
	//log.Info("TEST! ", resp.Services["c4bc3983-648c-45f9-a5b3-e1f6aa82db90"].PriceCents)
	//fmt.Printf("%#v", resp.Operators)
}

func testUpdateRead() {
	for {
		v := <-xmp_api_client.ChanCampaigns
		log.Info("UpdateTest: ", v.Id)
	}
}

func testAggregate() {
	var resp struct {
		Ok    bool   `json:"ok,omitempty"`
		Error string `json:"error,omitempty"`
	}

	items := make([]interface{}, 0)
	items = append(items, xmp_api_client.GetRandomAggregate())
	//items = append(items, xmp_api_client.GetRandomAggregate())

	err := xmp_api_client.Call("aggregate", &resp, items...)
	if err != nil {
		log.Error(err)
	}

	log.WithFields(
		log.Fields{
			"prefix": "testAggregate",
			"ok":     resp.Ok,
			"error":  resp.Error,
		},
	).Info("Request Done")
}
