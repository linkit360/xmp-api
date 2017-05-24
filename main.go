package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/xmp-api/src/client"
	"github.com/linkit360/xmp-api/src/server"
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
		DSN:        "go:40400",
		Timeout:    10,
		InstanceId: "2f4fd741-61ef-45ab-8436-840ce54d6d29",
	}

	if err := xmp_api_client.Init(cfg); err != nil {
		log.Panic("Cannot init acceptor client", err)
	}
	log.Info("Connected")

	/*
		var resp struct {
			Ok bool `json:"ok,omitempty"`
		}

		err := acceptorClient.Call("aggregate", &resp, acceptorClient.GetRandomAggregate(), acceptorClient.GetRandomAggregate())
		if err != nil {
			log.Error(err)
		}

		log.Debug(resp.Ok)
	*/

	/*
		var resp struct {
			Ok         bool `json:"ok,omitempty"`
			Status     int  `json:"status,omitempty"`
			OperatorId int  `json:"id_operator,omitempty"`
		}

		err := acceptorClient.Call("initialization", &resp)
		if err != nil {
			log.Error(err)
		}

		log.Debug(resp.Ok)
		log.Debug(resp.Status)
		log.Debug(resp.OperatorId)

		//resp, err := acceptor_client.SendAggregatedData(data2)
		//if err != nil {
		//	log.Errorln(err.Error())
		//}

		//log.Println(resp)

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
