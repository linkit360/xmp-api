package xmp_api_structs

type Country struct {
	Id       int64  `gorm:"primary_key",json:"id"` //
	Name     string `json:"name"`                  // "Russia"
	Code     int64  `json:"code"`                  // 7
	Iso      string `json:"iso"`                   // "RU"
	Flag     string `json:"flag"`                  // default "Unknown"
	Currency string `json:"currency,omitempty"`    //
}

// Tablenames for GORM
func (m Country) TableName() string {
	return "xmp_countries"
}
