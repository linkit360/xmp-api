package xmp_api_structs

type Service struct {
	Id              string    `gorm:"primary_key" json:"id"`         // UUID
	Code            string    `gorm:"column:id_service" json:"code"` // previous service id
	Title           string    `json:"title"`                         //
	Description     string    `json:"description,omitempty"`         //
	Price           int       `json:"price"`                         // Raw
	PriceCents      int       `json:"price_cents"`                   // In cents
	ContentIds      []string  `json:"content_ids,omitempty"`         // Unmarshalled content ids for use in content service (platform)
	Contents        []Content `json:"contents,omitempty"`            // Contents of service, unmarshalled and ready for use in platform
	ServiceOptsJson string    `gorm:"column:service_opts" json:"-"`  // taken from jsonb (database)
	ContentIdsJson  string    `gorm:"column:id_content" json:"-"`    // taken from jsonb (database)
	Status          int       `json:"status"`                        // Status, 1 = ok, 0 = deleted
	ProviderOpts
}

type ProviderOpts struct {
	// Comment format: Providers where it used (separated by comma) - comment
	SMSOnContent     string `json:"sms_on_content,omitempty"`       // Mobilink, QRTech - Текст СМС который отправляется при отправке контента. (FYI: может содержать URL и текст на разных языках)
	SMSOnSubscribe   string `json:"sms_on_subscribe,omitempty"`     // Mobilink - Текст СМС, который отправляется при подписке на сервис.
	SMSOnUnsubscribe string `json:"sms_on_unsubscribe,omitempty"`   // Mobilink - Текст СМС, который отправляется при отписке от сервиса.
	RetryDays        int    `json:"retry_days,omitempty"`           // Mobilink - days to keep retries, for periodic - subscription is alive
	InactiveDays     int    `json:"inactive_days,omitempty"`        // Mobilink - days of unsuccessful charge turns subscription into inactive state
	GraceDays        int    `json:"grace_days,omitempty"`           // Mobilink - days in end of subscription period where always must be charged OK
	PeriodicDays     string `json:"periodic_days,omitempty"`        // Mobilink - json - days of week to charge subscriber
	MinimalTouchTimes int    `json:"minimal_touch_times,omitempty"` // Mobilink - минимальное количество скачиваний контента пользователем в рамках одной подписки, чтобы подписка не была отменена.
	TrialDays         int    `json:"trial_days,omitempty"`          // Mobilink - Trial Days
	PurgeAfterDays    int    `json:"purge_after_days,omitempty"`    // Mobilink - Days before Purge policy

	// Other
	PaidHours           int    `json:"paid_hours,omitempty"`            // rejected rule
	DelayHours          int    `json:"delay_hours,omitempty"`           // repeat charge delay
	SMSOnCharged        string `json:"sms_on_charged,omitempty"`        //
	SMSOnRejected       string `json:"sms_on_rejected,omitempty"`       //
	SMSOnBlackListed    string `json:"sms_on_blacklisted,omitempty"`    //
	SMSOnPostPaid       string `json:"sms_on_postpaid,omitempty"`       //
	PeriodicAllowedFrom int    `json:"periodic_allowed_from,omitempty"` // send content in sms allowed from and to times.
	PeriodicAllowedTo   int    `json:"periodic_allowed_to,omitempty"`   //
}

type Content struct {
	Id    string `gorm:"primary_key" json:"id"`       // UUID (name of file in S3)
	Title string `json:"title"`                       // Title (for logs and humans)
	Name  string `gorm:"column:filename" json:"name"` // Name (filename inside zip)
}

// Tablenames for GORM
func (m Service) TableName() string {
	return "xmp_services"
}

func (m Content) TableName() string {
	return "xmp_content"
}
