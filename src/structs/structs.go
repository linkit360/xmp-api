package go_acceptor_structs

type HandShakeRequest struct {
	InstanceId string `json:"insance_id,omitempty"` // provider name by instance id
}

type HandShakeResponse struct {
	Ok        bool                `json:"ok"`                  // false if error, true if ok
	Error     string              `json:"error,omitempty"`     // error in case of any error
	BlackList []string            `json:"blacklist,omitempty"` // array of blacklisted numbers
	Campaigns map[string]Campaign `json:"campaigns,omitempty"` // map by uuid
	Services  map[string]Service  `json:"services,omitempty"`  // map by uuid
	Operators map[string]Operator `json:"operators,omitempty"` // map by anything
}

type Operator struct {
	Name        string // mobiilnk
	Code        int64  // mcc mnc: 41001
	CountryName string // pakistan
}

type AggregateResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type AggregateRequest struct {
	Aggregated []Aggregate `json:"aggregated,omitempty"`
}

type Aggregate struct {
	ProviderName           string `json:"provider_name,omitempty"`
	OperatorCode           int64  `json:"operator_code,omitempty"`
	ReportAt               int64  `json:"report_at,omitempty"`
	CampaignCode           string `json:"id_campaign,omitempty"`
	LpHits                 int64  `json:"lp_hits,omitempty"`
	LpMsisdnHits           int64  `json:"lp_msisdn_hits,omitempty"`
	MoTotal                int64  `json:"mo,omitempty"`
	MoChargeSuccess        int64  `json:"mo_charge_success,omitempty"`
	MoChargeSum            int64  `json:"mo_charge_sum,omitempty"`
	MoChargeFailed         int64  `json:"mo_charge_failed,omitempty"`
	MoRejected             int64  `json:"mo_rejected,omitempty"`
	Outflow                int64  `json:"outflow,omitempty"`
	RenewalTotal           int64  `json:"renewal,omitempty"`
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

type Service struct {
	Id                  string  `json:"id,omitempty"`            // unique id
	Code                string  `json:"code,omitempty"`          // previous service id
	Price               int     `json:"price,omitempty"`         // в целых рублях
	RetryDays           int     `json:"retry_days,omitempty"`    // for retries - days to keep retries, for periodic - subscription is alive
	InactiveDays        int     `json:"inactive_days,omitempty"` // days of unsuccessful charge turns subscription into inactive state
	GraceDays           int     `json:"grace_days,omitempty"`    // days in end of subscription period where always must be charged OK
	PaidHours           int     `json:"paid_hours,omitempty"`    // rejected rule
	DelayHours          int     `json:"delay_hours,omitempty"`   // repeat charge delay
	SMSOnSubscribe      string  `json:"sms_on_unsubscribe,omitempty"`
	SMSOnCharged        string  `json:"sms_on_charged,omitempty"`
	SMSOnUnsubscribe    string  `json:"sms_on_subscribe,omitempty"` // if empty, do not send
	SMSOnContent        string  `json:"sms_on_content,omitempty"`
	SMSOnRejected       string  `json:"sms_on_rejected,omitempty"`
	SMSOnBlackListed    string  `json:"sms_on_blacklisted,omitempty"`
	SMSOnPostPaid       string  `json:"sms_on_postpaid,omitempty"`
	PeriodicAllowedFrom int     `json:"periodic_allowed_from,omitempty"` // send content in sms allowed from and to times.
	PeriodicAllowedTo   int     `json:"periodic_allowed_to,omitempty"`
	PeriodicDays        string  `json:"periodic_days,omitempty"` // days of week to charge subscriber
	MinimalTouchTimes   int     `json:"minimal_touch_times,omitempty"`
	ContentIds          []int64 `json:"content_ids,omitempty"`
}

type Campaign struct {
	Id               string `json:"id,omitempty"`   // UUID
	Code             string `json:"code,omitempty"` // previous id
	Title            string `json:"title,omitempty"`
	Link             string `json:"link,omitempty"`
	Lp               string `json:"lp,omitempty"` // UUID
	Hash             string `json:"hash,omitempty"`
	ServiceCode      string `json:"service_code,omitempty"` // previous service code
	AutoClickRatio   int64  `json:"auto_click_ratio,omitempty"`
	AutoClickEnabled bool   `json:"auto_click_enabled,omitempty"`
	PageSuccess      string `json:"page_success,omitempty"`
	PageError        string `json:"page_error,omitempty"`
	PageThankYou     string `json:"page_thank_you,omitempty"`
	PageWelcome      string `json:"page_welcome,omitempty"`
}
