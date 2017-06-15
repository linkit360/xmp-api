package base

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func GetWsData() (map[string]uint64, map[string]string, uint64, uint64, uint64) {
	// widgets
	rows, err := db.Raw(
		"SELECT " +
			"SUM(lp_hits) AS LpHits, " +
			"SUM(mo_total) AS Mo, " +
			"SUM(mo_charge_success) AS MoSuccess " +
			"FROM xmp_reports WHERE " +
			"report_at >= '" + time.Now().Format("2006-01-02") + "'",
	).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var LpHits uint64 = 0
	var Mo uint64 = 0
	var MoSuccess uint64 = 0
	for rows.Next() {
		rows.Scan(
			&LpHits,
			&Mo,
			&MoSuccess,
		)
	}

	// map
	rows, err = db.Raw(
		"SELECT" +
			"(SELECT i.id FROM xmp_instances AS i WHERE p.id = i.id_provider) AS instance_id, " +
			"(SELECT c.iso FROM xmp_countries AS c WHERE c.id = p.id_country) AS country " +
			"FROM " +
			"xmp_providers AS p;",
	).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var iso string
	var instance string

	var provs = make(map[string]string)
	for rows.Next() {
		rows.Scan(
			&instance,
			&iso,
		)

		provs[instance] = iso
	}

	// TODO: providers by id_instance
	rows, err = db.Raw(
		"SELECT id_instance, SUM(lp_hits) " +
			"FROM xmp_reports " +
			"WHERE report_at >= '" + time.Now().Format("2006-01-02") + "' " +
			"GROUP BY id_instance",
	).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var prov string
	var sum uint64
	var countries = map[string]uint64{}
	for rows.Next() {
		rows.Scan(
			&prov,
			&sum,
		)

		if provs[prov] != "" {
			countries[provs[prov]] = countries[provs[prov]] + sum
		}
	}

	defer rows.Close()
	return countries, provs, LpHits, Mo, MoSuccess
}
