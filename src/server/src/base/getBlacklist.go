package base

import (
	"strconv"

	"github.com/linkit360/xmp-api/src/structs"
)

func GetBlacklist(id_provider int) ([]string, error) {
	//var err error
	out := make([]string, 0)
	data := make([]xmp_api_structs.Blacklist, 0)

	// Get campaigns by services ids
	db.Where("id_provider = ?", id_provider).Find(&data)
	for _, blacklist := range data {
		out = append(out, strconv.Itoa(blacklist.Msisdn))
	}

	return out, nil
}
