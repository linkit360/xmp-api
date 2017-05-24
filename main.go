package main

//log "github.com/Sirupsen/logrus"

//acceptorClient "github.com/linkit360/go-acceptor-client"

func main() {

	//for testing client && server

	/*
		cfg := acceptorClient.ClientConfig{
			Enabled:    true,
			DSN:        "go:50319",
			Timeout:    10,
			InstanceId: "2f4fd741-61ef-45ab-8436-840ce54d6d29",
		}

		if err := acceptorClient.Init(cfg); err != nil {
			log.Panic("Cannot init acceptor client", err)
		}
		log.Info("Connected")
	*/

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
	*/

	//resp, err := acceptor_client.SendAggregatedData(data2)
	//if err != nil {
	//	log.Errorln(err.Error())
	//}

	//log.Println(resp)

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
