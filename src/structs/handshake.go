package xmp_api_structs

type HandShake struct {
	Ok        bool                `json:"ok"`                  // false if error, true if ok
	Error     string              `json:"error,omitempty"`     // error in case of any error
	Country   Country             `json:"country,omitempty"`   // Country of Instance
	Services  map[string]Service  `json:"services,omitempty"`  // map by uuid
	Campaigns map[string]Campaign `json:"campaigns,omitempty"` // map by uuid
	Operators map[int64]Operator  `json:"operators,omitempty"` // map by anything
	BlackList []string            `json:"blacklist,omitempty"` // array of blacklisted numbers
	Pixels    []PixelSetting      `json:"pixel,omitempty"`     //
}
