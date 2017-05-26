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
		provOpts := xmp_api_structs.ServiceProvOpts{}
		err = json.Unmarshal([]byte(service.ServiceOpts), &provOpts)
		if err != nil {
			return nil, err
		}
		service.ProvOpts = append(service.ProvOpts, provOpts)
		service.ServiceOpts = ""

		// Content
		contentIds := make([]string, 0)
		err = json.Unmarshal([]byte(service.IdContent), &contentIds)
		if err != nil {
			return nil, err
		}
		service.IdContent = ""
		service.Content = make([]xmp_api_structs.Content, 0)

		db.Where("id IN (?)", contentIds).Find(&service.Content)

		// Append to return
		out[service.Id] = service
	}

	return out, nil
}
