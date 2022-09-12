package plivo

type CampaignService struct {
	client *Client
}
type CampaignCreationParams struct {
	BrandID            string    `json:"brand_id" url:"brand_id" validate:"required"`
	CampaignAlias      *string   `json:"campaign_alias,omitempty" url:"campaign_alias,omitempty"`
	Vertical           string    `json:"vertical" url:"vertical"`
	Usecase            string    `json:"usecase" url:"usecase"`
	ResellerID         string    `json:" reseller_id,omitempty" url:" reseller_id,omitempty"`
	SubUsecases        *[]string `json:"sub_usecases" url:"sub_usecases"`
	Description        string    `json:"description,omitempty" url:"description,omitempty"`
	EmbeddedLink       *bool     `json:"embedded_link,omitempty" url:"embedded_link,omitempty"`
	EmbeddedPhone      *bool     `json:"embedded_phone,omitempty" url:"embedded_phone,omitempty"`
	AgeGated           *bool     `json:"age_gated,omitempty" url:"age_gated,omitempty"`
	DirectLending      *bool     `json:"direct_lending,omitempty" url:"direct_lending,omitempty"`
	SubscriberOptin    bool      `json:"subscriber_optin" url:"subscriber_optin"`
	SubscriberOptout   bool      `json:"subscriber_optout" url:"subscriber_optout"`
	SubscriberHelp     bool      `json:"subscriber_help" url:"subscriber_help"`
	AffiliateMarketing bool      `json:"affiliate_marketing" url:"subscriber_help"`
	Sample1            *string   `json:"sample1" url:"sample1"`
	Sample2            *string   `json:"sample2,omitempty" url:"sample2,omitempty"`
	URL                string    `json:"url,omitempty" url:"url,omitempty"`
	Method             string    `json:"method,omitempty" url:"method,omitempty"`
	MessageFlow        string    `json:"message_flow,omitempty" url:"message_flow"`
	HelpMessage        string    `json:"help_message,omitempty" url:"help_message"`
	OptinKeywords      string    `json:"optin_keywords,omitempty" url:"optin_keywords"`
	OptinMessage       string    `json:"optin_message,omitempty" url:"optin_message"`
	OptoutKeywords     string    `json:"optout_keywords,omitempty" url:"optout_keywords"`
	OptoutMessage      string    `json:"optout_message,omitempty" url:"optout_message"`
	HelpKeywords       string    `json:"help_keywords,omitempty" url:"help_keywords"`
}

type CampaignListResponse struct {
	ApiID            string     `json:"api_id,omitempty"`
	CampaignResponse []Campaign `json:"campaigns,omitempty"`
}

type CampaignGetResponse struct {
	ApiID    string   `json:"api_id"`
	Campaign Campaign `json:"campaign"`
}

type CampaignCreateResponse struct {
	ApiID      string `json:"api_id,omitempty"`
	CampaignID string `json:"campaign_id"`
	Message    string `json:"message,omitempty"`
}

type Campaign struct {
	BrandID            string      `json:"brand_id,omitempty"`
	CampaignID         string      `json:"campaign_id,omitempty"`
	MnoMetadata        MnoMetadata `json:"mno_metadata,omitempty"`
	ResellerID         string      `json:"reseller_id,omitempty"`
	Usecase            string      `json:"usecase,omitempty"`
	SubSsecase         string      `json:"sub_usecase,omitempty"`
	RegistrationStatus string      `json:"registration_status,omitempty"`
	MessageFlow        string      `json:"message_flow,omitempty"`
	HelpMessage        string      `json:"help_message,omitempty"`
	OptinKeywords      string      `json:"optin_keywords,omitempty"`
	OptinMessage       string      `json:"optin_message,omitempty"`
	OptoutKeywords     string      `json:"optout_keywords,omitempty"`
	OptoutMessage      string      `json:"optout_message,omitempty"`
	HelpKeywords       string      `json:"help_keywords,omitempty"`
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
	Limit   int     `url:"limit,omitempty"`
	Offset  int     `url:"offset,omitempty"`
}

type CampaignNumberLinkParams struct {
	Numbers []string `json:"numbers,omitempty"`
	URL     string   `json:"url,omitempty"`
	Method  string   `json:"method,omitempty"`
}

type CampaignNumberUnlinkParams struct {
	URL    string `json:"url,omitempty"`
	Method string `json:"method,omitempty"`
}

type CampaignNumbersGetParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type CampaignNumberLinkUnlinkResponse struct {
	ApiID   string `json:"api_id"`
	Message string `json:"message,omitempty"`
}

type CampaignNumberGetResponse struct {
	ApiID           string           `json:"api_id"`
	CampaignID      string           `json:"campaign_id"`
	CampaignAlias   string           `json:"campaign_alias"`
	CampaignUseCase string           `json:"usecase"`
	CampaignNumbers []CampaignNumber `json:"phone_numbers"`
}

type CampaignNumber struct {
	Number string `json:"number"`
	Status string `json:"status"`
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

func (service *CampaignService) Get(campaignID string) (response *CampaignGetResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "10dlc/Campaign/%s", campaignID)
	if err != nil {
		return
	}
	response = &CampaignGetResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) Create(params CampaignCreationParams) (response *CampaignCreateResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "10dlc/Campaign")
	if err != nil {
		return
	}
	response = &CampaignCreateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) NumberLink(campaignID string, params CampaignNumberLinkParams) (response *CampaignNumberLinkUnlinkResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "10dlc/Campaign/%s/Number", campaignID)
	if err != nil {
		return
	}
	response = &CampaignNumberLinkUnlinkResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) NumberGet(campaignID, number string) (response *CampaignNumberGetResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "10dlc/Campaign/%s/Number/%s", campaignID, number)
	if err != nil {
		return
	}
	response = &CampaignNumberGetResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) NumbersGet(campaignID string, params CampaignNumbersGetParams) (response *CampaignNumberGetResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "10dlc/Campaign/%s/Number", campaignID)
	if err != nil {
		return
	}
	response = &CampaignNumberGetResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *CampaignService) NumberUnlink(campaignID, number string, params CampaignNumberUnlinkParams) (response *CampaignNumberLinkUnlinkResponse, err error) {
	req, err := service.client.NewRequest("DELETE", params, "10dlc/Campaign/%s/Number/%s", campaignID, number)
	if err != nil {
		return
	}
	response = &CampaignNumberLinkUnlinkResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}
