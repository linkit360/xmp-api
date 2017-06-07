package xmp_api_structs

type Operator struct {
	Id          int64  `gorm:"primary_key",json:"id"`  //
	Code        int64  `json:"code"`                   // mcc mnc: 41001
	Name        string `json:"name,omitempty"`         // mobiilnk
	CountryName string `json:"country_name,omitempty"` // pakistan
}

// Tablenames for GORM
func (m Operator) TableName() string {
	return "xmp_operators"
}
