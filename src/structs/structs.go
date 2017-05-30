package xmp_api_structs

type HandShake struct {
	Ok        bool                `json:"ok"`                  // false if error, true if ok
	Error     string              `json:"error,omitempty"`     // error in case of any error
	Status    int                 `json:"status,omitempty"`    // Instance status
	Services  map[string]Service  `json:"services,omitempty"`  // map by uuid
	Campaigns map[string]Campaign `json:"campaigns,omitempty"` // map by uuid
	Operators map[int64]Operator  `json:"operators,omitempty"` // map by anything
	BlackList []string            `json:"blacklist,omitempty"` // array of blacklisted numbers
	Pixels    []PixelSetting      `json:"pixel,omitempty"`     //
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
	Provider               string `json:"provider_name,omitempty"`
	OperatorCode           int64  `json:"operator_code,omitempty"`
	ReportAt               int64  `json:"report_at,omitempty"`
	CampaignCode           string `json:"campaign_code,omitempty"`
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

type Campaign struct {
	Id               string `json:"id,omitempty"`                 // UUID
	Code             string `json:"code,omitempty"`               // previous id
	Title            string `json:"title,omitempty"`              //
	Link             string `json:"link,omitempty"`               //
	Lp               string `json:"lp,omitempty"`                 // UUID
	Hash             string `json:"hash,omitempty"`               //
	ServiceCode      string `json:"service_code,omitempty"`       // previous service code
	AutoClickRatio   int64  `json:"auto_click_ratio,omitempty"`   //
	AutoClickEnabled bool   `json:"auto_click_enabled,omitempty"` //
	PageSuccess      string `json:"page_success,omitempty"`       //
	PageError        string `json:"page_error,omitempty"`         //
	PageThankYou     string `json:"page_thank_you,omitempty"`     //
	PageWelcome      string `json:"page_welcome,omitempty"`       //
}

type PixelSetting struct {
	Id           string
	CampaignCode string
	OperatorCode int64
	Publisher    string
	Endpoint     string
	Timeout      int
	Enabled      bool
	Ratio        int
}
