package base

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/linkit360/xmp-api/src/structs"
)

func GetServices(id_provider int) (out map[string]xmp_api_structs.Service, err error) {
	out = make(map[string]xmp_api_structs.Service)

	defer func() {
		if err != nil {
			log.Errorln("SVC ERR: " + err.Error())
		}
	}()

	data := make([]xmp_api_structs.Service, 0)

	// Get services by provider id
	db.Debug().Where("status = 1 AND id_provider = ?", id_provider).Find(&data)

	log.Info("SVC 1: " + strconv.Itoa(len(data)))

	for _, service := range data {

		log.Info("SVC 2: " + service.Id)

		// Provider specific options
		provOpts := xmp_api_structs.ProviderOpts{}
		err = json.Unmarshal([]byte(service.ServiceOptsJson), &provOpts)
		if err != nil {
			return
		}
		service.ServiceOptsJson = ""

		var provOptsTmp []byte
		provOptsTmp, err = json.Marshal(provOpts)
		if err != nil {
			return
		}

		log.Info("SVC 2.1: " + string(provOptsTmp))

		err = json.Unmarshal(provOptsTmp, &service.ProviderOpts)
		if err != nil {
			return
		}

		// Content
		contentIds := make([]string, 0)
		err = json.Unmarshal([]byte(service.ContentIdsJson), &contentIds)
		if err != nil {
			return
		}

		log.Info("SVC 2.2: " + strconv.Itoa(len(contentIds)))

		service.ContentIdsJson = ""
		service.Contents = make([]xmp_api_structs.Content, 0)

		db.Where("status = 1 AND id IN (?)", contentIds).Find(&service.Contents)

		log.Info("SVC 3: " + strconv.Itoa(len(service.Contents)))

		// Append to return
		out[service.Id] = service
	}

	log.Info("SVC 4: " + strconv.Itoa(len(out)))

	return
}
