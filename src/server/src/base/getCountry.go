package base

import (
	"errors"
	"log"

	"github.com/linkit360/xmp-api/src/structs"
)

func GetCountry(id_provider int) (xmp_api_structs.Country, error) {
	var err error
	out := xmp_api_structs.Country{}

	rows, err := db.Raw("SELECT id_country FROM xmp_providers WHERE id = ? LIMIT 1;", id_provider).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var id_country int64
	for rows.Next() {
		rows.Scan(
			&id_country,
		)
	}

	if id_country == 0 {
		return out, errors.New("No country")
	}

	// Get country
	db.Where("status = 1 AND id = ?", id_country).Find(&out)

	return out, nil
}
