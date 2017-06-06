package base

import (
	"github.com/linkit360/xmp-api/src/structs"
)

func GetCampaigns(services map[string]xmp_api_structs.Service) (map[string]xmp_api_structs.Campaign, error) {
	//var err error
	out := make(map[string]xmp_api_structs.Campaign)
	data := make([]xmp_api_structs.Campaign, 0)

	serviceIds := make([]string, 0)
	for _, service := range services {
		serviceIds = append(serviceIds, service.Id)
	}

	// Get campaigns by services ids
	db.Where("id_service in (?)", serviceIds).Find(&data)
	for _, campaign := range data {
		out[campaign.Id] = campaign
	}

	return out, nil
}
