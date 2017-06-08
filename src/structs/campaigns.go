package xmp_api_structs

type Campaign struct {
	Id               string `gorm:"primary_key",json:"id"`        // UUID
	Code             string `gorm:"column:id_old",json:"code"`    // previous id
	Title            string `json:"title,omitempty"`              //
	Link             string `json:"link"`                         //
	Lp               string `json:"lp"`                           // UUID
	Hash             string `json:"hash,omitempty"`               //
	ServiceCode      string `json:"service_code,omitempty"`       // previous service code
	AutoClickRatio   int64  `json:"auto_click_ratio,omitempty"`   //
	AutoClickEnabled bool   `json:"auto_click_enabled,omitempty"` //
	PageSuccess      string `json:"page_success,omitempty"`       //
	PageError        string `json:"page_error,omitempty"`         //
	PageThankYou     string `json:"page_thank_you,omitempty"`     //
	PageWelcome      string `json:"page_welcome,omitempty"`       //
}

// Tablenames for GORM
func (m Campaign) TableName() string {
	return "xmp_campaigns"
}
