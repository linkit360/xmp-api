package base

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/linkit360/xmp-api/src/server/src/config"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/jinzhu/gorm.v1"
)

var cfg config.DbConfig
var db *gorm.DB

func Init() {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(log.DebugLevel)
	var err error

	cfg = config.LoadConfig()
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.Host,
			cfg.User,
			cfg.Database,
			cfg.Password,
		),
	)

	if err != nil {
		log.Panic("Base: Connection Error ", err.Error())
	}
}

func GetOptions(instanceId string) (int, int) {
	rows, err := db.Raw("SELECT status,id_operator FROM xmp_instances WHERE id = ?", instanceId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var status int
	var operatorId int
	for rows.Next() {
		rows.Scan(
			&status,
			&operatorId,
		)
	}

	log.Debug(status)
	log.Debug(operatorId)

	return status, operatorId
}

/*
func SaveRows(rows []xmp_api_structs.Aggregate) error {
	var query string = fmt.Sprintf(
		"INSERT INTO %sreports ("+

			"report_at, "+
			"provider_name, "+
			"operator_code, "+
			"id_campaign, "+
			"lp_hits, "+
			"lp_msisdn_hits, "+

			"mo_total, "+
			"mo_charge_success, "+
			"mo_charge_sum, "+
			"mo_charge_failed, "+
			"mo_rejected, "+

			"outflow, "+
			"renewal_total, "+
			"renewal_charge_success, "+
			"renewal_charge_sum, "+
			"renewal_failed, "+

			"pixels,"+

			"injection_total, "+
			"injection_charge_success, "+
			"injection_charge_sum, "+
			"injection_failed, "+

			"expired_total, "+
			"expired_charge_success, "+
			"expired_charge_sum, "+
			"expired_failed"+

			") VALUES ("+

			"to_timestamp($1), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25"+

			");",
		config.TablePrefix)

	//TODO: make bulk request
	for _, row := range rows {
		if _, err := pgsql.Exec(
			query,
			row.ReportAt,
			row.ProviderName,
			row.OperatorCode,
			row.CampaignCode,
			row.LpHits,
			row.LpMsisdnHits,

			row.MoTotal,
			row.MoChargeSuccess,
			row.MoChargeSum,
			row.MoChargeFailed,
			row.MoRejected,

			row.Outflow,
			row.RenewalTotal,
			row.RenewalChargeSuccess,
			row.RenewalChargeSum,
			row.RenewalFailed,

			row.Pixels,

			row.InjectionTotal,
			row.InjectionChargeSuccess,
			row.InjectionChargeSum,
			row.InjectionFailed,

			row.ExpiredTotal,
			row.ExpiredChargeSuccess,
			row.ExpiredChargeSum,
			row.ExpiredFailed,
		); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func GetWsData() (map[string]uint64, map[string]string, uint64, uint64, uint64) {
	// widgets
	rows, err := pgsql.Query("SELECT " +
		"SUM(lp_hits) AS LpHits, " +
		"SUM(mo_total) AS Mo, " +
		"SUM(mo_charge_success) AS MoSuccess " +
		"FROM xmp_reports WHERE " +
		"report_at >= '" + time.Now().Format("2006-01-02") + "'",
	)
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
	//noinspection SqlResolve
	rows, err = pgsql.Query("SELECT " +
		"c.iso, p.name, p.id " +
		"FROM xmp_providers as p " +
		"INNER JOIN xmp_countries as c " +
		"ON (p.id_country = c.id);",
	)
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

	//fmt.Printf("%+v", provs)

	//noinspection SqlResolve
	rows, err = pgsql.Query("SELECT provider_name, SUM(lp_hits) FROM xmp_reports WHERE report_at >= '" + time.Now().Format("2006-01-02") + "' GROUP BY provider_name")
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
	//fmt.Printf("%+v", countries)

	return countries, provs, LpHits, Mo, MoSuccess
}
*/

/*
func GetBlackList(providerName string, time string) []string {
	//noinspection SqlResolve
	rows, err := pgsql.Query("SELECT id FROM xmp_providers WHERE name = '" + providerName + "'")
	if err != nil {
		log.Fatal(err)
	}

	var id string
	for rows.Next() {
		rows.Scan(
			&id,
		)
	}
	//fmt.Printf("%+v", id)

	var data []string
	if id != "" {
		query := "SELECT msisdn FROM xmp_msisdn_blacklist WHERE id_provider = " + id
		if time != "" {
			query = query + " AND created_at >= '" + time + "'"
		}

		//noinspection SqlResolve
		rows, err = pgsql.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		var msisdn string
		for rows.Next() {
			rows.Scan(
				&msisdn,
			)
			data = append(data, msisdn)
		}
	}

	return data
}

func GetCampaigns(provider string) []acceptorStructs.CampaignsCampaign {
	rows, err := pgsql.Query("SELECT id FROM xmp_providers WHERE name_alias = '" + provider + "';")
	if err != nil {
		log.Fatal(err)
	}

	var id uint
	for rows.Next() {
		rows.Scan(
			&id,
		)
	}
	//log.Infoln(id)

	data := make([]acceptorStructs.CampaignsCampaign, 0)
	if id > 0 {
		var query = fmt.Sprintf("SELECT id FROM xmp_operators WHERE id_provider = %d;", id)
		//log.Infoln(query)

		rows, err := pgsql.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		ids := make([]uint, 0)
		var id uint
		for rows.Next() {
			rows.Scan(
				&id,
			)
			ids = append(ids, id)
		}

		//log.Infoln(ids)

		if len(ids) > 0 {
			query = "SELECT id, title, link, id_lp FROM xmp_campaigns WHERE id_operator IN(0"

			for _, value := range ids {
				query = query + fmt.Sprintf(", %d", value)
			}

			query = query + ");"
			//log.Infoln(query)

			rows, err := pgsql.Query(query)
			if err != nil {
				log.Fatal(err)
			}

			var camp acceptorStructs.CampaignsCampaign
			for rows.Next() {
				rows.Scan(
					&camp.Id,
					&camp.Title,
					&camp.Link,
					&camp.Lp,
				)

				//log.Infoln(camp)

				data = append(data, camp)
			}
		}
	}

	//log.Infoln(data)
	return data
}
*/
