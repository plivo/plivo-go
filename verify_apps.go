package plivo

type VerifyAppService struct {
	client *Client
	VerifyApp
}

type VerifyAppCreateParams struct {
	Name                  string `json:"name" url:"name"`
	OtpType               string `json:"otp_type,omitempty" url:"otp_type,omitempty"`
	OtpLength             int    `json:"otp_length,omitempty" url:"otp_length,omitempty"`
	OtpExpiry             int    `json:"otp_expiry,omitempty" url:"otp_expiry,omitempty"`
	OtpAttempts           int    `json:"otp_attempts,omitempty" url:"otp_attempts,omitempty"`
	BrandName             string `json:"brand_name,omitempty" url:"brand_name,omitempty"`
	SmsChannel            *bool  `json:"sms_channel,omitempty" url:"sms_channel,omitempty"`
	VoiceChannel          *bool  `json:"voice_channel,omitempty" url:"voice_channel,omitempty"`
	WaChannel             *bool  `json:"wa_channel,omitempty" url:"wa_channel,omitempty"`
	IsDefault             *bool  `json:"is_default,omitempty" url:"is_default,omitempty"`
	TemplateUUID          string `json:"template_uuid,omitempty" url:"template_uuid,omitempty"`
	MessageRedaction      *bool  `json:"message_redaction,omitempty" url:"message_redaction,omitempty"`
	CustomerAppHash       string `json:"customer_app_hash,omitempty" url:"customer_app_hash,omitempty"`
	MaxValidationAttempts int    `json:"max_validation_attempts,omitempty" url:"max_validation_attempts,omitempty"`
	EnableFraudshield     *bool  `json:"enable_fraudshield,omitempty" url:"enable_fraudshield,omitempty"`
	FsProtectionLevel     string `json:"fs_protection_level,omitempty" url:"fs_protection_level,omitempty"`
	WabaID                string `json:"waba_id,omitempty" url:"waba_id,omitempty"`
	WabaPhoneNumber       string `json:"waba_phone_number,omitempty" url:"waba_phone_number,omitempty"`
	WabaTemplateID        string `json:"waba_template_id,omitempty" url:"waba_template_id,omitempty"`
}

type VerifyAppUpdateParams struct {
	Name                  string `json:"name,omitempty" url:"name,omitempty"`
	BrandName             string `json:"brand_name,omitempty" url:"brand_name,omitempty"`
	OtpType               string `json:"otp_type,omitempty" url:"otp_type,omitempty"`
	OtpLength             int    `json:"otp_length,omitempty" url:"otp_length,omitempty"`
	OtpExpiry             int    `json:"otp_expiry,omitempty" url:"otp_expiry,omitempty"`
	OtpAttempts           int    `json:"otp_attempts,omitempty" url:"otp_attempts,omitempty"`
	SmsChannel            *bool  `json:"sms_channel,omitempty" url:"sms_channel,omitempty"`
	VoiceChannel          *bool  `json:"voice_channel,omitempty" url:"voice_channel,omitempty"`
	WaChannel             *bool  `json:"wa_channel,omitempty" url:"wa_channel,omitempty"`
	IsDefault             *bool  `json:"is_default,omitempty" url:"is_default,omitempty"`
	TemplateUUID          string `json:"template_uuid,omitempty" url:"template_uuid,omitempty"`
	MessageRedaction      *bool  `json:"message_redaction,omitempty" url:"message_redaction,omitempty"`
	CustomerAppHash       string `json:"customer_app_hash,omitempty" url:"customer_app_hash,omitempty"`
	MaxValidationAttempts int    `json:"max_validation_attempts,omitempty" url:"max_validation_attempts,omitempty"`
	EnableFraudshield     *bool  `json:"enable_fraudshield,omitempty" url:"enable_fraudshield,omitempty"`
	FsProtectionLevel     string `json:"fs_protection_level,omitempty" url:"fs_protection_level,omitempty"`
	WabaID                string `json:"waba_id,omitempty" url:"waba_id,omitempty"`
	WabaPhoneNumber       string `json:"waba_phone_number,omitempty" url:"waba_phone_number,omitempty"`
	WabaTemplateID        string `json:"waba_template_id,omitempty" url:"waba_template_id,omitempty"`
}

type VerifyAppListParams struct {
	Name               string `url:"name,omitempty"`
	AppUUID            string `url:"app_uuid,omitempty"`
	Channel            string `url:"channel,omitempty"`
	Status             string `url:"status,omitempty"`
	Limit              int    `url:"limit,omitempty"`
	Offset             int    `url:"offset,omitempty"`
	CreatedAt          string `url:"created_at,omitempty"`
	CreatedAtLt        string `url:"created_at__lt,omitempty"`
	CreatedAtLte       string `url:"created_at__lte,omitempty"`
	CreatedAtGt        string `url:"created_at__gt,omitempty"`
	CreatedAtGte       string `url:"created_at__gte,omitempty"`
	SubaccountAuthID   string `url:"subaccount_auth_id,omitempty"`
}

type VerifyApp struct {
	AppUUID               string `json:"app_uuid,omitempty"`
	Name                  string `json:"name,omitempty"`
	OtpType               string `json:"otp_type,omitempty"`
	OtpLength             int    `json:"otp_length,omitempty"`
	OtpExpiry             int    `json:"otp_expiry,omitempty"`
	OtpAttempts           int    `json:"otp_attempts,omitempty"`
	BrandName             string `json:"brand_name,omitempty"`
	SmsChannel            *bool  `json:"sms_channel,omitempty"`
	VoiceChannel          *bool  `json:"voice_channel,omitempty"`
	WaChannel             *bool  `json:"wa_channel,omitempty"`
	IsDefault             *bool  `json:"is_default,omitempty"`
	TemplateUUID          string `json:"template_uuid,omitempty"`
	MessageRedaction      *bool  `json:"message_redaction,omitempty"`
	CustomerAppHash       string `json:"customer_app_hash,omitempty"`
	MaxValidationAttempts int    `json:"max_validation_attempts,omitempty"`
	EnableFraudshield     *bool  `json:"enable_fraudshield,omitempty"`
	FsProtectionLevel     string `json:"fs_protection_level,omitempty"`
	WabaID                string `json:"waba_id,omitempty"`
	WabaPhoneNumber       string `json:"waba_phone_number,omitempty"`
	WabaTemplateID        string `json:"waba_template_id,omitempty"`
}

type VerifyWhatsapp struct {
	WabaID          string `json:"waba_id,omitempty"`
	WabaPhoneNumber string `json:"waba_phone_number,omitempty"`
	WabaTemplateID  string `json:"waba_template_id,omitempty"`
}

type VerifyAppCreateResponseBody struct {
	ApiID   string `json:"api_id,omitempty"`
	AppUUID string `json:"app_uuid,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type VerifyAppUpdateResponseBody struct {
	ApiID   string `json:"api_id,omitempty"`
	AppUUID string `json:"app_uuid,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type VerifyAppDeleteResponseBody struct {
	ApiID   string `json:"api_id,omitempty"`
	AppUUID string `json:"app_uuid,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type VerifyAppGetResponseBody struct {
	ApiID          string          `json:"api_id,omitempty"`
	VerifyApp      *VerifyApp      `json:"verify_app,omitempty"`
	VerifyWhatsapp *VerifyWhatsapp `json:"verify_whatsapp,omitempty"`
}

type VerifyAppListMeta struct {
	Previous   *string `json:"previous,omitempty"`
	Next       *string `json:"next,omitempty"`
	Limit      int64   `json:"limit,omitempty"`
	Offset     int64   `json:"offset,omitempty"`
	TotalCount int64   `json:"total_count,omitempty"`
}

type VerifyAppList struct {
	ApiID       string            `json:"api_id,omitempty"`
	Meta        VerifyAppListMeta `json:"meta"`
	VerifyApps  []VerifyApp       `json:"verify_apps"`
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

func (service *VerifyAppService) Delete(appUUID string) (response *VerifyAppDeleteResponseBody, err error) {
	req, err := service.client.NewRequest("DELETE", nil, "Verify/App/%s", appUUID)
	if err != nil {
		return
	}
	response = &VerifyAppDeleteResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}