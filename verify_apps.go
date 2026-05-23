package plivo

type VerifyAppService struct {
	client *Client
	VerifyApp
}

type VerifyAppCreateParams struct {
	Name                  string `json:"name" url:"name"`
	BrandName             string `json:"brand_name,omitempty" url:"brand_name,omitempty"`
	OtpType               string `json:"otp_type,omitempty" url:"otp_type,omitempty"`
	OtpLength             int    `json:"otp_length,omitempty" url:"otp_length,omitempty"`
	OtpExpiry             int    `json:"otp_expiry,omitempty" url:"otp_expiry,omitempty"`
	OtpAttempts           int    `json:"otp_attempts,omitempty" url:"otp_attempts,omitempty"`
	MaxValidationAttempts int    `json:"max_validation_attempts,omitempty" url:"max_validation_attempts,omitempty"`
	SmsChannel            *bool  `json:"sms_channel,omitempty" url:"sms_channel,omitempty"`
	VoiceChannel          *bool  `json:"voice_channel,omitempty" url:"voice_channel,omitempty"`
	WaChannel             *bool  `json:"wa_channel,omitempty" url:"wa_channel,omitempty"`
	WabaId                string `json:"waba_id,omitempty" url:"waba_id,omitempty"`
	WabaPhoneNumber       string `json:"waba_phone_number,omitempty" url:"waba_phone_number,omitempty"`
	WabaTemplateId        string `json:"waba_template_id,omitempty" url:"waba_template_id,omitempty"`
	TemplateUUID          string `json:"template_uuid,omitempty" url:"template_uuid,omitempty"`
	IsDefault             *bool  `json:"is_default,omitempty" url:"is_default,omitempty"`
	MessageRedaction      *bool  `json:"message_redaction,omitempty" url:"message_redaction,omitempty"`
	EnableFraudshield     *bool  `json:"enable_fraudshield,omitempty" url:"enable_fraudshield,omitempty"`
	FsProtectionLevel     string `json:"fs_protection_level,omitempty" url:"fs_protection_level,omitempty"`
	CustomerAppHash       string `json:"customer_app_hash,omitempty" url:"customer_app_hash,omitempty"`
	NumberPool            string `json:"number_pool,omitempty" url:"number_pool,omitempty"`
}

type VerifyAppCreateResponseBody struct {
	ApiID   string `json:"api_id"`
	AppUUID string `json:"app_uuid"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type VerifyAppListParams struct {
	Name           string `url:"name,omitempty"`
	Subaccount     string `url:"subaccount,omitempty"`
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
	CreatedAt      string `url:"created_at,omitempty"`
	CreatedAtLt    string `url:"created_at__lt,omitempty"`
	CreatedAtLte   string `url:"created_at__lte,omitempty"`
	CreatedAtGt    string `url:"created_at__gt,omitempty"`
	CreatedAtGte   string `url:"created_at__gte,omitempty"`
}

type VerifyAppMeta struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Offset   int64   `json:"offset"`
	Limit    int64   `json:"limit"`
}

type VerifyApp struct {
	ApiID             string `json:"api_id,omitempty"`
	AppUUID           string `json:"app_uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	BrandName         string `json:"brand_name,omitempty"`
	OtpType           string `json:"otp_type,omitempty"`
	OtpLength         int    `json:"otp_length,omitempty"`
	OtpExpiry         int    `json:"otp_expiry,omitempty"`
	OtpAttempts       int    `json:"otp_attempts,omitempty"`
	MaxValidationAttempts int `json:"max_validation_attempts,omitempty"`
	SmsChannel        *bool  `json:"sms_channel,omitempty"`
	VoiceChannel      *bool  `json:"voice_channel,omitempty"`
	WaChannel         *bool  `json:"wa_channel,omitempty"`
	WabaId            string `json:"waba_id,omitempty"`
	WabaPhoneNumber   string `json:"waba_phone_number,omitempty"`
	WabaTemplateId    string `json:"waba_template_id,omitempty"`
	TemplateUUID      string `json:"template_uuid,omitempty"`
	IsDefault         *bool  `json:"is_default,omitempty"`
	MessageRedaction  *bool  `json:"message_redaction,omitempty"`
	EnableFraudshield *bool  `json:"enable_fraudshield,omitempty"`
	FsProtectionLevel string `json:"fs_protection_level,omitempty"`
	CustomerAppHash   string `json:"customer_app_hash,omitempty"`
	NumberPool        string `json:"number_pool,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
}

type VerifyAppWhatsApp struct {
	WabaId          string `json:"waba_id,omitempty"`
	WabaPhoneNumber string `json:"waba_phone_number,omitempty"`
	WabaTemplateId  string `json:"waba_template_id,omitempty"`
}

type VerifyAppList struct {
	ApiID      string      `json:"api_id"`
	VerifyApps []VerifyApp `json:"verify_apps"`
	Meta       VerifyAppMeta `json:"meta"`
}

type VerifyAppGetResponseBody struct {
	ApiID         string             `json:"api_id"`
	VerifyApp     *VerifyApp         `json:"verify_app"`
	VerifyWhatsApp *VerifyAppWhatsApp `json:"verify_whatsapp"`
}

type VerifyAppUpdateParams struct {
	Name                  string `json:"name,omitempty" url:"name,omitempty"`
	BrandName             string `json:"brand_name,omitempty" url:"brand_name,omitempty"`
	OtpType               string `json:"otp_type,omitempty" url:"otp_type,omitempty"`
	OtpLength             int    `json:"otp_length,omitempty" url:"otp_length,omitempty"`
	OtpExpiry             int    `json:"otp_expiry,omitempty" url:"otp_expiry,omitempty"`
	OtpAttempts           int    `json:"otp_attempts,omitempty" url:"otp_attempts,omitempty"`
	MaxValidationAttempts int    `json:"max_validation_attempts,omitempty" url:"max_validation_attempts,omitempty"`
	SmsChannel            *bool  `json:"sms_channel,omitempty" url:"sms_channel,omitempty"`
	VoiceChannel          *bool  `json:"voice_channel,omitempty" url:"voice_channel,omitempty"`
	WaChannel             *bool  `json:"wa_channel,omitempty" url:"wa_channel,omitempty"`
	WabaId                string `json:"waba_id,omitempty" url:"waba_id,omitempty"`
	WabaPhoneNumber       string `json:"waba_phone_number,omitempty" url:"waba_phone_number,omitempty"`
	WabaTemplateId        string `json:"waba_template_id,omitempty" url:"waba_template_id,omitempty"`
	TemplateUUID          string `json:"template_uuid,omitempty" url:"template_uuid,omitempty"`
	IsDefault             *bool  `json:"is_default,omitempty" url:"is_default,omitempty"`
	MessageRedaction      *bool  `json:"message_redaction,omitempty" url:"message_redaction,omitempty"`
	EnableFraudshield     *bool  `json:"enable_fraudshield,omitempty" url:"enable_fraudshield,omitempty"`
	FsProtectionLevel     string `json:"fs_protection_level,omitempty" url:"fs_protection_level,omitempty"`
	CustomerAppHash       string `json:"customer_app_hash,omitempty" url:"customer_app_hash,omitempty"`
	Client                string `json:"client,omitempty" url:"client,omitempty"`
}

type VerifyAppUpdateResponseBody struct {
	ApiID   string `json:"api_id"`
	AppUUID string `json:"app_uuid"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type VerifyAppTemplate struct {
	TemplateUUID string `json:"template_uuid,omitempty"`
	Name         string `json:"name,omitempty"`
	Body         string `json:"body,omitempty"`
	Language     string `json:"language,omitempty"`
}

type VerifyAppListTemplatesResponseBody struct {
	ApiID     string              `json:"api_id"`
	Templates []VerifyAppTemplate `json:"templates"`
}

func (service *VerifyAppService) Create(params VerifyAppCreateParams) (response *VerifyAppCreateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Verify/App")
	if err != nil {
		return
	}
	response = &VerifyAppCreateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyAppService) List(params VerifyAppListParams) (response *VerifyAppList, err error) {
	req, err := service.client.NewRequest("GET", params, "Verify/App")
	if err != nil {
		return
	}
	response = &VerifyAppList{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyAppService) ListTemplates() (response *VerifyAppListTemplatesResponseBody, err error) {
	req, err := service.client.NewRequest("GET", nil, "Verify/App/templates")
	if err != nil {
		return
	}
	response = &VerifyAppListTemplatesResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyAppService) Get(appUUID string) (response *VerifyAppGetResponseBody, err error) {
	req, err := service.client.NewRequest("GET", nil, "Verify/App/%s", appUUID)
	if err != nil {
		return
	}
	response = &VerifyAppGetResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyAppService) Update(appUUID string, params VerifyAppUpdateParams) (response *VerifyAppUpdateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Verify/App/%s", appUUID)
	if err != nil {
		return
	}
	response = &VerifyAppUpdateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyAppService) Delete(appUUID string) (err error) {
	req, err := service.client.NewRequest("DELETE", nil, "Verify/App/%s", appUUID)
	if err != nil {
		return
	}
	err = service.client.ExecuteRequest(req, nil)
	return
}