package xmp_api_structs

type Campaign struct {
	Id               string `gorm:"primary_key" json:"id"`                              // UUID
	Code             string `gorm:"column:id_old" json:"id_old"`                        // previous id
	Title            string `json:"title,omitempty"`                                    //
	Link             string `json:"link"`                                               //
	Lp               string `gorm:"column:id_lp" json:"lp"`                             // UUID
	Hash             string `json:"hash,omitempty"`                                     //
	ServiceId        string `gorm:"column:id_service" json:"id_service"`                // service ID / UUID
	ServiceCode      string `json:"service_code"`                                       //
	AutoClickRatio   int64  `gorm:"column:autoclick_ratio" json:"auto_click_ratio"`     //
	AutoClickEnabled bool   `gorm:"column:autoclick_enabled" json:"auto_click_enabled"` //
	PageSuccess      string `json:"page_success,omitempty"`                             //
	PageError        string `json:"page_error,omitempty"`                               //
	PageThankYou     string `json:"page_thank_you,omitempty"`                           //
	PageWelcome      string `json:"page_welcome,omitempty"`                             //
	Status           int    `json:"status"`                                             // Status, 1 = ok, 0 = deleted
}

// Tablenames for GORM
func (m Campaign) TableName() string {
	return "xmp_campaigns"
}
