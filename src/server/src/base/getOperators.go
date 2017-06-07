package base

import (
	"github.com/linkit360/xmp-api/src/structs"
)

func GetOperators(id_provider int) (map[int64]xmp_api_structs.Operator, error) {
	//var err error
	out := make(map[int64]xmp_api_structs.Operator)
	data := make([]xmp_api_structs.Operator, 0)

	// Get operators by provider id
	db.Where("status = 1 AND id_provider = ?", id_provider).Find(&data)
	for _, operator := range data {
		out[operator.Id] = operator
	}

	return out, nil
}
