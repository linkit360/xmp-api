package xmp_api_structs

type HandShake struct {
	Ok        bool                `json:"ok"`                  // false if error, true if ok
	Error     string              `json:"error,omitempty"`     // error in case of any error
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
