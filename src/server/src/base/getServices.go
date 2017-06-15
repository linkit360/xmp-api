package base

import (
	"encoding/json"

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
	db.Where("status = 1 AND id_provider = ?", id_provider).Find(&data)
	for _, service := range data {
		// Provider specific options
		if service.ServiceOptsJson != "" && service.ServiceOptsJson != "{}" && service.ServiceOptsJson != "[]" {
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

			err = json.Unmarshal(provOptsTmp, &service.ProviderOpts)
			if err != nil {
				return
			}
		}

		// Content
		contentIds := make([]string, 0)
		err = json.Unmarshal([]byte(service.ContentIdsJson), &contentIds)
		if err != nil {
			return
		}

		service.ContentIdsJson = ""
		service.Contents = make([]xmp_api_structs.Content, 0)

		db.Where("status = 1 AND id IN (?)", contentIds).Find(&service.Contents)

		// Append to return
		out[service.Id] = service
	}

	return
}
