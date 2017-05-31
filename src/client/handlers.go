package xmp_api_client

import (
	"math/rand"
	"time"

	xmp_api_structs "github.com/linkit360/xmp-api/src/structs"
)

/*
func SendAggregatedData(data []acceptorStructs.Aggregate) (acceptorStructs.AggregateResponse, error) {
	var res acceptorStructs.AggregateResponse
	err := call(
		"Aggregate.Receive",
		acceptorStructs.AggregateRequest{Aggregated: data},
		&res,
	)
	return res, err
}
*/

func GetRandomAggregate() xmp_api_structs.Aggregate {
	return xmp_api_structs.Aggregate{
		InstanceId:           "a7da1e9f-fcc1-4087-9c58-4d31bcdbd515",
		ReportAt:             time.Now().UTC().Unix(),
		CampaignCode:         "290",
		OperatorCode:         52000,
		LpHits:               rand.Int63n(200),
		LpMsisdnHits:         rand.Int63n(150),
		MoTotal:              rand.Int63n(200),
		MoChargeSuccess:      rand.Int63n(200),
		MoChargeSum:          1000.,
		MoChargeFailed:       rand.Int63n(200),
		MoRejected:           rand.Int63n(200),
		RenewalTotal:         rand.Int63n(200),
		RenewalChargeSuccess: rand.Int63n(200),
		RenewalChargeSum:     12312.,
		RenewalFailed:        rand.Int63n(200),
		Pixels:               rand.Int63n(200),
	}
}
