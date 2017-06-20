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

	// MAP
	// Get countries id-iso mapping
	var countr = make(map[int64]string)
	rows, err = db.Raw("SELECT id,iso from xmp_countries;").Rows()
	if err != nil {
		log.Fatal(err)
	}
	var id int64
	var iso string
	for rows.Next() {
		rows.Scan(
			&id,
			&iso,
		)

		countr[id] = iso
	}

	// Get id_country - id_instance mapping
	rows, err = db.Raw(
		"SELECT " +
			"(SELECT p.id_country FROM xmp_providers AS p WHERE p.id = i.id_provider LIMIT 1) AS id_country, " +
			"id as id_instance " +
			"FROM xmp_instances AS i;",
	).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var id_country int64
	var id_instance string
	var provs = make(map[string]string)
	for rows.Next() {
		rows.Scan(
			&id_country,
			&id_instance,
		)

		provs[id_instance] = countr[id_country]
	}

	// Get LP Hits by id_instance (for today)
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
		} else {
			log.Error("Empty provs for: ", prov)
		}
	}

	defer rows.Close()
	return countries, provs, LpHits, Mo, MoSuccess
}
