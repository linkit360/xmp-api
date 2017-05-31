package base

import (
	"fmt"

	"github.com/linkit360/xmp-api/src/structs"
)

func SaveRows(rows []xmp_api_structs.Aggregate) error {
	//TODO: make bulk request
	for _, row := range rows {
		if _, err := db.Raw(
			"INSERT INTO xmp_reports ("+

				"report_at, "+
				"id_instance, "+
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
			row.ReportAt,
			row.InstanceId,
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
		).Rows(); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
