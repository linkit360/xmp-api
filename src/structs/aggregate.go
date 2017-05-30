package xmp_api_structs

type AggregateResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type AggregateRequest struct {
	Aggregated []Aggregate `json:"aggregated,omitempty"`
}

type Aggregate struct {
	Provider        string `json:"provider"`
	OperatorCode    int64  `json:"operator_code"`
	ReportAt        int64  `json:"report_at"`
	CampaignCode    string `json:"campaign_code"`
	LpHits          int64  `json:"lp_hits,omitempty"`
	LpMsisdnHits    int64  `json:"lp_msisdn_hits,omitempty"`
	MoTotal         int64  `json:"mo,omitempty"`
	MoChargeSuccess int64  `json:"mo_charge_success,omitempty"`
	MoChargeSum     int64  `json:"mo_charge_sum,omitempty"`
	MoChargeFailed  int64  `json:"mo_charge_failed,omitempty"`
	MoRejected      int64  `json:"mo_rejected,omitempty"`
	Outflow         int64  `json:"outflow,omitempty"`
	RenewalTotal    int64  `json:"renewal,omitempty"`
	RenewalChargeSuccess   int64  `json:"renewal_charge_success,omitempty"`
	RenewalChargeSum       int64  `json:"renewal_charge_sum,omitempty"`
	RenewalFailed          int64  `json:"renewal_failed,omitempty"`
	InjectionTotal         int64  `json:"injection,omitempty"`
	InjectionChargeSuccess int64  `json:"injection_charge_success,omitempty"`
	InjectionChargeSum     int64  `json:"injection_charge_sum,omitempty"`
	InjectionFailed        int64  `json:"injection_failed,omitempty"`
	ExpiredTotal           int64  `json:"expired,omitempty"`
	ExpiredChargeSuccess   int64  `json:"expired_charge_success,omitempty"`
	ExpiredChargeSum       int64  `json:"expired_charge_sum,omitempty"`
	ExpiredFailed          int64  `json:"expired_failed,omitempty"`
	Pixels                 int64  `json:"pixels,omitempty"`
}
