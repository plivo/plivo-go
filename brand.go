package plivo

type BrandService struct {
	client *Client
}

type BrandCreationParams struct {
	AltBusinessIDType  *string `json:"alt_business_id_type,omitempty" url:"alt_business_id_type,omitempty"`
	AltBusinessID      *string `json:"alt_business_id,omitempty" url:"alt_business_id,omitempty"`
	City               string  `json:"city,omitempty" url:"city,omitempty"`
	CompanyName        string  `json:"company_name,omitempty" url:"company_name,omitempty"`
	Country            string  `json:"country,omitempty" url:"country,omitempty"`
	Ein                string  `json:"ein,omitempty" url:"ein,omitempty"`
	EinIssuingCountry  string  `json:"ein_issuing_country,omitempty" url:"ein_issuing_country,omitempty"`
	Email              string  `json:"email,omitempty" url:"email,omitempty"`
	EntityType         string  `json:"entity_type,omitempty" url:"entity_type,omitempty"`
	FirstName          *string `json:"first_name,omitempty" url:"first_name,omitempty"`
	LastName           *string `json:"last_name,omitempty" url:"last_name,omitempty"`
	Phone              string  `json:"phone,omitempty" url:"phone,omitempty"`
	PostalCode         string  `json:"postal_code,omitempty" url:"postal_code,omitempty"`
	RegistrationStatus string  `json:"registration_status,omitempty" url:"registration_status,omitempty"`
	State              string  `json:"state,omitempty" url:"state,omitempty"`
	StockExchange      string  `json:"stock_exchange,omitempty" url:"stock_exchange,omitempty"`
	StockSymbol        string  `json:"stock_symbol,omitempty" url:"stock_symbol,omitempty"`
	Street             string  `json:"street,omitempty" url:"street,omitempty"`
	Vertical           string  `json:"vertical,omitempty" url:"vertical,omitempty"`
	Website            *string `json:"website,omitempty" url:"website,omitempty"`
	SecondaryVetting   *string `json:"secondary_vetting,omitempty" url:"secondary_vetting,omitempty"`
}
type BrandCreationResponse struct {
	ApiID string        `json:"api_id,omitempty"`
	Brand BrandResponse `json:"brand,omitempty"`
}
type BrandResponse struct {
	AltBusinessIDType  string `json:"alt_business_id_type,omitempty"`
	BrandID            string `json:"brand_id,omitempty"`
	City               string `json:"city,omitempty"`
	Country            string `json:"country,omitempty"`
	EinIssuingCountry  string `json:"ein_issuing_country,omitempty"`
	Email              string `json:"email,omitempty"`
	EntityType         string `json:"entity_type,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	Phone              string `json:"phone,omitempty"`
	PostalCode         string `json:"postal_code,omitempty"`
	RegistrationStatus string `json:"registration_status,omitempty"`
	State              string `json:"state,omitempty"`
	StockExchange      string `json:"stock_exchange,omitempty"`
	StockSymbol        string `json:"stock_symbol,omitempty"`
	Vertical           string `json:"vertical,omitempty"`
	Website            string `json:"website,omitempty"`
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
	BrandID            string `json:"brand_id,omitempty"`
	CompanyName        string `json:"company_name,omitempty"`
	Ein                string `json:"ein,omitempty"`
	EinIssuingCountry  string `json:"ein_issuing_country,omitempty"`
	Email              string `json:"email,omitempty"`
	EntityType         string `json:"entity_type,omitempty"`
	RegistrationStatus string `json:"registration_status,omitempty"`
	Vertical           string `json:"vertical,omitempty"`
	Website            string `json:"website,omitempty"`
}
type BrandListParams struct {
	Type   *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
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
