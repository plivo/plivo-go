package plivo

type BrandService struct {
	client *Client
}

type BrandCreationParams struct {
	BrandAlias       string  `json:"brand_alias" url:"brand_alias" validate:"required"`
	Type             string  `json:"brand_type" url:"brand_type" validate:"oneof= STARTER STANDARD ''"`
	ProfileUUID      string  `json:"profile_uuid" url:"profile_uuid" validate:"required,max=36"`
	SecondaryVetting *string `json:"secondary_vetting,omitempty" url:"secondary_vetting,omitempty"`
	URL              string  `json:"url,omitempty" url:"url,omitempty"`
	Method           string  `json:"method,omitempty" url:"method,omitempty"`
	SubAccountId     string  `json:"subaccount_id,omitempty" url:"subaccount_id,omitempty"`
	DefaultCampaignCreationRequest
}

type DefaultCampaignCreationRequest struct {
	EmailRecipients     string   `json:"emailRecipients,omitempty" url:"emailRecipients,omitempty"`
	CampaignName        string   `json:"campaignName,omitempty" url:"campaignName,omitempty"`
	CampaignUseCase     string   `json:"campaignUseCase,omitempty" url:"campaignUseCase,omitempty"`
	CampaignSubUseCases []string `json:"campaignSubUseCases,omitempty" url:"campaignSubUseCases,omitempty"`
	CampaignDescription string   `json:"campaignDescription,omitempty" url:"campaignDescription,omitempty"`
	SampleMessage1      string   `json:"sampleMessage1,omitempty" url:"sampleMessage1,omitempty"`
	SampleMessage2      string   `json:"sampleMessage2,omitempty" url:"sampleMessage2,omitempty"`
	EmbeddedLink        bool     `json:"embeddedLink,omitempty" url:"embeddedLink,omitempty"`
	EmbeddedPhone       bool     `json:"embeddedPhone,omitempty" url:"embeddedPhone,omitempty"`
	NumberPool          bool     `json:"numberPool,omitempty" url:"numberPool,omitempty"`
	AgeGated            bool     `json:"ageGated,omitempty" url:"ageGated,omitempty"`
	DirectLending       bool     `json:"directLending,omitempty" url:"directLending,omitempty"`
	SubscriberOptin     bool     `json:"subscriberOptin,omitempty" url:"subscriberOptin,omitempty"`
	SubscriberOptout    bool     `json:"subscriberOptout,omitempty" url:"subscriberOptout,omitempty"`
	SubscriberHelp      bool     `json:"subscriberHelp,omitempty" url:"subscriberHelp,omitempty"`
	AffiliateMarketing  bool     `json:"affiliateMarketing,omitempty" url:"affiliateMarketing,omitempty"`
	ResellerID          string   `json:"resellerID,omitempty" url:"resellerID,omitempty"`
}
type BrandCreationResponse struct {
	ApiID   string `json:"api_id,omitempty"`
	BrandID string `json:"brand_id,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type BrandListResponse struct {
	ApiID         string  `json:"api_id,omitempty"`
	BrandResponse []Brand `json:"brands,omitempty"`
}

type BrandGetResponse struct {
	ApiID string `json:"api_id,omitempty"`
	Brand Brand  `json:"brand,omitempty"`
}
type Brand struct {
	BrandAlias         string            `json:"brand_alias,omitempty"`
	EntityType         string            `json:"entity_type,omitempty"`
	BrandID            string            `json:"brand_id,omitempty"`
	ProfileUUID        string            `json:"profile_uuid,omitempty"`
	FirstName          string            `json:"first_name,omitempty"`
	LastName           string            `json:"last_name,omitempty"`
	Name               string            `json:"name,omitempty"`
	CompanyName        string            `json:"company_name,omitempty"`
	BrandType          string            `json:"brand_type,omitempty"`
	Ein                string            `json:"ein,omitempty"`
	EinIssuingCountry  string            `json:"ein_issuing_country,omitempty"`
	StockSymbol        string            `json:"stock_symbol,omitempty"`
	StockExchange      string            `json:"stock_exchange,omitempty"`
	Website            string            `json:"website,omitempty"`
	Vertical           string            `json:"vertical,omitempty"`
	AltBusinessID      string            `json:"alt_business_id,omitempty"`
	AltBusinessidType  string            `json:"alt_business_id_type,omitempty"`
	RegistrationStatus string            `json:"registration_status,omitempty"`
	VettingStatus      string            `json:"vetting_status,omitempty"`
	VettingScore       int64             `json:"vetting_score,omitempty"`
	Address            Address           `json:"address,omitempty"`
	AuthorizedContact  AuthorizedContact `json:"authorized_contact,omitempty"`
}
type BrandListParams struct {
	Type   *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
}

type Address struct {
	Street     string `json:"street" validate:"max=100"`
	City       string `json:"city" validate:"max=100"`
	State      string `json:"state" validate:"max=20"`
	PostalCode string `json:"postal_code" validate:"max=10"`
	Country    string `json:"country" validate:"max=2"`
}

type AuthorizedContact struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty" validate:"max=16"`
	Email     string `json:"email,omitempty" validate:"max=100"`
	Title     string `json:"title,omitempty"`
	Seniority string `json:"seniority,omitempty"`
}

func (service *BrandService) List(params BrandListParams) (response *BrandListResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "10dlc/Brand")
	if err != nil {
		return
	}
	response = &BrandListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *BrandService) Get(brandID string) (response *BrandGetResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "10dlc/Brand/%s", brandID)
	if err != nil {
		return
	}
	response = &BrandGetResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *BrandService) Create(params BrandCreationParams) (response *BrandCreationResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "10dlc/Brand")
	if err != nil {
		return
	}
	response = &BrandCreationResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}
