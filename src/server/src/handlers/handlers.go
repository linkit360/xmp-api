package handlers

import (
	log "github.com/sirupsen/logrus"
)

var Clients map[string]string

func init() {
	Clients = make(map[string]string)
}

func Send(instance_id string, payload []byte) {
	if Clients[instance_id] == "" {
		log.Error("Send: No Instance")
		return
	}

	var resp struct {
		Message string `json:"message,omitempty"`
	}

	err := Call("update", Clients[instance_id], &resp, payload)
	if err != nil {
		log.Error("Send Error: ", resp.Message)
		return
	}

	log.Debug("Send: ", resp.Message)
}

/*
type BlackList struct{}

func (rpc *BlackList) GetAll(req acceptorStructs.BlackListGetParams, res *acceptorStructs.BlackListResponse) error {
	log.WithFields(log.Fields{
		"prefix":       "Handlers",
		"ProviderName": req.ProviderName,
	}).Info("BL GetAll")

	res.Msisdns = base.GetBlackList(req.ProviderName, "")

	//log.Printf("%+v\n", list)

	return nil
}

func (rpc *BlackList) GetNew(req acceptorStructs.BlackListGetParams, res *acceptorStructs.BlackListResponse) error {
	log.WithFields(log.Fields{
		"prefix":       "Handlers",
		"ProviderName": req.ProviderName,
		"Time":         req.Time,
	}).Info("BL GetNew")

	res.Msisdns = base.GetBlackList(req.ProviderName, req.Time)

	//log.Printf("%+v\n", list)

	return nil
}
*/
