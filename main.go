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
		Enabled: true,
		DSN:     "go:50318",
		Timeout: 10,
		//InstanceId: "a7da1e9f-fcc1-4087-9c58-4d31bcdbd515", // qrtech
		InstanceId: "58fbedf7-1abc-402b-8c2a-89fe256d32d9", // mobilink
	}

	if err := xmp_api_client.Init(cfg); err != nil {
		log.Panic("Cannot init acceptor client", err)
	}
	log.Info("Ready")

	testInitialization()
	//testAggregate()
	//testUpdateRead()
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

	//log.Info("TEST! ", resp.Campaigns["373bfcb3-f967-4860-96da-39637856f67b"].Lp)
	//fmt.Printf("%#v", resp.Operators)
}

func testUpdateRead() {
	for {
		v := <-xmp_api_client.ChanCampaigns
		log.Info("UpdateTest: ", v.Code)
	}
}
