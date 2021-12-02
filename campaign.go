package plivo

type CampaignService struct {
	client *Client
}
type CampaignCreationParams struct {
	BrandID          string    `json:"brand_id,omitempty" url:"brand_id,omitempty"`
	CampaignAlias    *string   `json:"campaign_alias,omitempty" url:"campaign_alias,omitempty"`
	Vertical         string    `json:"vertical,omitempty" url:"vertical,omitempty"`
	Usecase          string    `json:"usecase,omitempty" url:"usecase,omitempty"`
	SubUsecases      *[]string `json:"sub_usecases,omitempty" url:"sub_usecases,omitempty"`
	Description      string    `json:"description,omitempty" url:"description,omitempty"`
	EmbeddedLink     *bool     `json:"embedded_link,omitempty" url:"embedded_link,omitempty"`
	EmbeddedPhone    *bool     `json:"embedded_phone,omitempty" url:"embedded_phone,omitempty"`
	AgeGated         *bool     `json:"age_gated,omitempty" url:"age_gated,omitempty"`
	DirectLending    *bool     `json:"direct_lending,omitempty" url:"direct_lending,omitempty"`
	SubscriberOptin  bool      `json:"subscriber_optin,omitempty" url:"subscriber_optin,omitempty"`
	SubscriberOptout bool      `json:"subscriber_optout,omitempty" url:"subscriber_optout,omitempty"`
	SubscriberHelp   bool      `json:"subscriber_help,omitempty" url:"subscriber_help,omitempty"`
	Sample1          *string   `json:"sample1,omitempty" url:"sample1,omitempty"`
	Sample2          *string   `json:"sample2,omitempty" url:"sample2"`
}

type CampaignListResponse struct {
	ApiID            string     `json:"api_id,omitempty"`
	CampaignResponse []Campaign `json:"campaigns,omitempty"`
}

type CampaignResponse struct {
	ApiID    string   `json:"api_id,omitempty"`
	Campaign Campaign `json:"campaign,omitempty"`
}

type Campaign struct {
	BrandID     string      `json:"brand_id,omitempty"`
	CampaignID  string      `json:"campaign_id,omitempty"`
	MnoMetadata MnoMetadata `json:"mno_metadata,omitempty"`
	ResellerID  string      `json:"reseller_id,omitempty"`
	Usecase     string      `json:"usecase,omitempty"`
}
type MnoMetadata struct {
	ATandT          OperatorDetail `json:"AT&T,omitempty"`
	TMobile         OperatorDetail `json:"T-Mobile,omitempty"`
	VerizonWireless OperatorDetail `json:"Verizon Wireless,omitempty"`
	USCellular      OperatorDetail `json:"US Cellular,omitempty"`
}
type OperatorDetail struct {
	BrandTier string `json:"brand_tier,omitempty"`
	TPM       int    `json:"tpm,omitempty"`
}
type CampaignListParams struct {
	BrandID *string `json:"brand,omitempty"`
	Usecase *string `json:"usecase,omitempty"`
}

func (service *CampaignService) List(params CampaignListParams) (response *CampaignListResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "10dlc/Campaign")
	if err != nil {
		return
	}
	response = &CampaignListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) Get(campaignID string) (response *CampaignResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "10dlc/Campaign/%s", campaignID)
	if err != nil {
		return
	}
	response = &CampaignResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) Create(params CampaignCreationParams) (response *CampaignResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "10dlc/Campaign")
	if err != nil {
		return
	}
	response = &CampaignResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}
