package base

import (
	log "github.com/sirupsen/logrus"
)

func GetOptions(instanceId string) (int, int) {
	rows, err := db.Raw("SELECT status,id_provider FROM xmp_instances WHERE id = ?", instanceId).Rows()
	if err != nil {
		log.Error(err)
		return 0, 0
	}
	defer rows.Close()

	var status int
	var id_provider int
	for rows.Next() {
		rows.Scan(
			&status,
			&id_provider,
		)
	}

	return status, id_provider
}
