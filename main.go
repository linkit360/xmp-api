package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/client"
	"github.com/linkit360/xmp-api/src/server"
	"github.com/linkit360/xmp-api/src/structs"
	"github.com/x-cray/logrus-prefixed-formatter"
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
		Timeout:    10,
		InstanceId: "a7da1e9f-fcc1-4087-9c58-4d31bcdbd515",
	}

	if err := xmp_api_client.Init(cfg); err != nil {
		log.Panic("Cannot init acceptor client", err)
	}
	log.Info("Ready")

	//for {
	// instance doing his work (wait)
	testInitialization()
	//time.Sleep(3 * time.Second)
	//}
	//testAggregate()

	testUpdateRead()

	//for {
	//	time.Sleep(3 * time.Second)
	//}

	/*
		// Get BL All
		data, err := acceptorClient.GetBlackListed("cheese")
		if err != nil {
			log.Println("Error")
			log.Fatalln(err.Error())
		}

		log.Println("DATA")
		log.Printf("%+v\n", data)

		// Get BL Time
		data, err = acceptor_client.BlackListGetNew("cheese", "2000-01-01")
		if err != nil {
			log.Println("Error")
			log.Fatalln(err.Error())
		}

		log.Println("DATA")
		log.Printf("%+v\n", data)

		// Send Aggregate
		data2 := []acceptor.Aggregate{
			acceptor_client.GetRandomAggregate(),
		}

		//log.Println(data)
		resp, err := acceptor_client.SendAggregatedData(data2)
		if err != nil {
			log.Println("Error")
			log.Println(err.Error())
		}

		log.Println(resp)

		data, err := acceptor_client.CampaignsGet("cheese")
		if err != nil {
			log.Println("Error")
			log.Fatalln(err.Error())
		}

		log.Println("DATA")
		log.Printf("%+v\n", data)
	*/
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

	log.Debug(resp.Ok)
	log.Debug(resp.Error)
}

func testInitialization() {
	var resp xmp_api_structs.HandShake
	err := xmp_api_client.Call("initialization", &resp)
	if err != nil {
		log.Error(err)
	}

	log.WithFields(
		log.Fields{
			"prefix":   "testInitialization",
			"ok":       resp.Ok,
			"error":    resp.Error,
			"services": len(resp.Services),
		},
	).Info("Request Done")

	//log.Printf("%#v\n", resp.Services["6f257e12-f1f1-47d4-9a43-5bb966f94d6a"])

	log.Info("")
}

func testUpdateRead() {
	for {
		v := <-xmp_api_client.ChanServices
		log.Info("UpdateTest: ", v.Id)
	}
}
