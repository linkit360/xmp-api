package base

import (
	"time"

	log "github.com/Sirupsen/logrus"
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
	defer rows.Close()

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
		"SELECT " +
			"c.iso, p.name, p.id " +
			"FROM xmp_providers as p " +
			"INNER JOIN xmp_countries as c " +
			"ON (p.id_country = c.id);",
	).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var iso string
	var prov string
	var id uint64

	var provs = make(map[string]string)
	for rows.Next() {
		rows.Scan(
			&iso,
			&prov,
			&id,
		)

		provs[prov] = iso
	}

	// TODO: providers by instance_id
	rows, err = db.Raw(
		"SELECT instance_id, SUM(lp_hits) " +
			"FROM xmp_reports " +
			"WHERE report_at >= '" + time.Now().Format("2006-01-02") + "' " +
			"GROUP BY instance_id",
	).Rows()
	if err != nil {
		log.Fatal(err)
	}

	var sum uint64
	var countries = map[string]uint64{}
	for rows.Next() {
		rows.Scan(
			&prov,
			&sum,
		)

		countries[provs[prov]] = sum
	}

	return countries, provs, LpHits, Mo, MoSuccess
}
