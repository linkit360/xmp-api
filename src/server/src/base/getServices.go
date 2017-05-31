package base

import (
	"encoding/json"

	"github.com/linkit360/xmp-api/src/structs"
)

func GetServices(id_provider int) (map[string]xmp_api_structs.Service, error) {
	var err error
	out := make(map[string]xmp_api_structs.Service)
	data := make([]xmp_api_structs.Service, 0)

	// Get services by provider id
	db.Where("id_provider = ?", id_provider).Find(&data)
	for _, service := range data {
		// Provider specific options
		provOpts := xmp_api_structs.ProviderOpts{}
		err = json.Unmarshal([]byte(service.ServiceOptsJson+"sdf"), &provOpts)
		if err != nil {
			return nil, err
		}
		service.ServiceOptsJson = ""

		provOptsTmp, err := json.Marshal(provOpts)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(provOptsTmp, &service.ProviderOpts)
		if err != nil {
			return nil, err
		}

		// Content
		contentIds := make([]string, 0)
		err = json.Unmarshal([]byte(service.ContentIdsJson), &contentIds)
		if err != nil {
			return nil, err
		}
		service.ContentIdsJson = ""
		service.Contents = make([]xmp_api_structs.Content, 0)

		db.Where("id IN (?)", contentIds).Find(&service.Contents)

		// Append to return
		out[service.Id] = service
	}

	return out, nil
}
