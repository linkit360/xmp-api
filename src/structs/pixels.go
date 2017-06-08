package xmp_api_structs

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
