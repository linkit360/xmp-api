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

	/*
		// AWS S3
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIAILRGTUB6EBNVUPFA",
				"8Hf/b4jldspVA2hCUlBqAJhmpCjr7M1zAU/LYjrl",
				"",
			),
		})
		if err != nil {
			log.Error("Cannot init S3 Session ", err)
		}

		svc := s3.New(sess)

		ctx := context.Background()

		result, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
			Bucket: aws.String("xmp-lp"),
			Key:    aws.String("31788279-604c-416f-868c-c96dbdbe3e3c"),
		})
		if err != nil {
			log.Error("Cannot init S3 Client ", err)
		}

		defer result.Body.Close()

		log.Info("Object Etag: ", aws.StringValue(result.ETag))

		return
	*/

	//for {
	// instance doing his work (wait)
	testInitialization()
	//time.Sleep(3 * time.Second)
	//}
	//testAggregate()

	//testUpdateRead()

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
			"prefix":   "testInitialization",
			"ok":       resp.Ok,
			"error":    resp.Error,
			"services": len(resp.Services),
		},
	).Info("Request Done")

	//log.Info(resp.Operators[2])
}

func testUpdateRead() {
	for {
		v := <-xmp_api_client.ChanServices
		log.Info("UpdateTest: ", v.Id)
	}
}
