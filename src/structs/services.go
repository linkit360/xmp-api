package xmp_api_structs

type Service struct {
	Id              string    `gorm:"primary_key",json:"id,omitempty"` // UUID
	Code            string    `json:"code,omitempty"`                  // previous service id
	Title           string    `json:"title,omitempty"`                 //
	Description     string    `json:"description,omitempty"`           //
	Price           int       `json:"price,omitempty"`                 // In cents
	ContentIds      []string  `json:"content_ids,omitempty"`           // Unmarshalled content ids for use in content service (platform)
	Contents        []Content `json:"contents,omitempty"`              // Contents of service, unmarshalled and ready for use in platform
	ServiceOptsJson string    `gorm:"column:service_opts",json:"-"`    // taken from jsonb (database)
	ContentIdsJson  string    `gorm:"column:id_content",json:"-"`      // taken from jsonb (database)
	ProviderOpts
}

type ProviderOpts struct {
	SMSOnContent string `json:"sms_on_content,omitempty"` // QRTech

	// Other
	RetryDays           int    `json:"retry_days,omitempty"`            // for retries - days to keep retries, for periodic - subscription is alive
	InactiveDays        int    `json:"inactive_days,omitempty"`         // days of unsuccessful charge turns subscription into inactive state
	GraceDays           int    `json:"grace_days,omitempty"`            // days in end of subscription period where always must be charged OK
	PaidHours           int    `json:"paid_hours,omitempty"`            // rejected rule
	DelayHours          int    `json:"delay_hours,omitempty"`           // repeat charge delay
	SMSOnSubscribe      string `json:"sms_on_unsubscribe,omitempty"`    //
	SMSOnCharged        string `json:"sms_on_charged,omitempty"`        //
	SMSOnUnsubscribe    string `json:"sms_on_subscribe,omitempty"`      // if empty, do not send
	SMSOnRejected       string `json:"sms_on_rejected,omitempty"`       //
	SMSOnBlackListed    string `json:"sms_on_blacklisted,omitempty"`    //
	SMSOnPostPaid       string `json:"sms_on_postpaid,omitempty"`       //
	PeriodicAllowedFrom int    `json:"periodic_allowed_from,omitempty"` // send content in sms allowed from and to times.
	PeriodicAllowedTo   int    `json:"periodic_allowed_to,omitempty"`   //
	PeriodicDays        string `json:"periodic_days,omitempty"`         // days of week to charge subscriber
	MinimalTouchTimes   int    `json:"minimal_touch_times,omitempty"`   //
}

type Content struct {
	Id    string `gorm:"primary_key",json:"id"` // UUID
	Title string `json:"title"`                 // Title
	Name  string `json:"name"`                  // Name
}

// Tablenames for GORM
func (m Service) TableName() string {
	return "xmp_services"
}

func (m Content) TableName() string {
	return "xmp_content"
}
