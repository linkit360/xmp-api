package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Call(funcName string, addr string, res interface{}, request []byte) error {
	var url string = "http://" + addr + "/" + funcName
	var err error

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	/*
		// GET by default
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		if len(req) > 0 {
			// POST
			jsonValue, err := json.Marshal(req)
			if err != nil {
				return err
			}
	*/

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	if err != nil {
		return err
	}
	//}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	return nil
}
