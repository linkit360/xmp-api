package xmp_api_structs

type Blacklist struct {
	Id         string `gorm:"primary_key" json:"id"` // UUID
	OperatorId int64  `json:"id_operator"`           // operator of number
	Msisdn     int    `json:"msisdn"`                // phone number
	Status     int    `json:"status"`                // Status, 1 = ok, 0 = deleted
}

// Tablenames for GORM
func (m Blacklist) TableName() string {
	return "xmp_msisdn_blacklist"
}
