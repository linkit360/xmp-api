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
	rows, err := db.Raw("SELECT status,id_provider FROM xmp_instances WHERE id = ?", instanceId).Rows()
	if err != nil {
		log.Fatal(err)
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
