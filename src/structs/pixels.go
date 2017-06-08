package xmp_api_structs

type PixelSetting struct {
	Id           string `json:"id,omitempty"`
	CampaignCode string `json:"campaign_code,omitempty"`
	OperatorCode int64  `json:"operator_code,omitempty"`
	Publisher    string `json:"publisher,omitempty"`
	Endpoint     string `json:"endpoint,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	Enabled      bool   `json:"enabled,omitempty"`
	Ratio        int    `json:"ratio,omitempty"`
}
