package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ── Requirements ─────────────────────────────────────────────────────────────

func TestPhoneNumberComplianceRequirementService_Get(t *testing.T) {
	expectResponse("phoneNumberComplianceRequirementGetResponse.json", 200)

	resp, err := client.PhoneNumberComplianceRequirements.Get(PhoneNumberComplianceRequirementGetParams{
		CountryISO: "IN",
		NumberType: "local",
		UserType:   "business",
	})
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assert.Equal(t, "1b592b77-d072-4298-b193-370ba887f94a", resp.RequirementID)
	assert.Equal(t, "IN", resp.CountryISO)
	assert.Equal(t, "local", resp.NumberType)
	assert.Equal(t, "business", resp.UserType)
	assert.Equal(t, 2, len(resp.DocumentTypes))
	assert.Equal(t, "Registration Certificate", resp.DocumentTypes[0].Name)
	assert.Equal(t, true, resp.DocumentTypes[0].ProofRequired)
	assert.Equal(t, 1, len(resp.DocumentTypes[0].RequiredFields))
	assert.Equal(t, "business_name", resp.DocumentTypes[0].RequiredFields[0].FieldName)

	assertRequest(t, "GET", "PhoneNumber/Compliance/Requirements")
}

func TestPhoneNumberComplianceRequirementService_Get_NilClient(t *testing.T) {
	expectResponse("phoneNumberComplianceRequirementGetResponse.json", 200)

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumberComplianceRequirements.Get(PhoneNumberComplianceRequirementGetParams{
		CountryISO: "IN",
		NumberType: "local",
		UserType:   "business",
	})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

// ── Compliance Get ───────────────────────────────────────────────────────────

func TestPhoneNumberComplianceService_Get(t *testing.T) {
	expectResponse("phoneNumberComplianceGetResponse.json", 200)
	complianceId := "767bc62c-2332-4a34-959c-1f6416186254"

	resp, err := client.PhoneNumberCompliance.Get(complianceId, PhoneNumberComplianceGetParams{})
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assert.Equal(t, complianceId, resp.Compliance.ComplianceID)
	assert.Equal(t, "TestBiz-100ac06f", resp.Compliance.Alias)
	assert.Equal(t, "submitted", resp.Compliance.Status)
	assert.Equal(t, "IN", resp.Compliance.CountryISO)
	assert.Equal(t, "local", resp.Compliance.NumberType)
	assert.Equal(t, "business", resp.Compliance.UserType)
	assert.Equal(t, "https://example.com/webhook", resp.Compliance.CallbackURL)
	assert.Equal(t, "POST", resp.Compliance.CallbackMethod)

	assertRequest(t, "GET", "PhoneNumber/Compliance/%s", complianceId)
}

func TestPhoneNumberComplianceService_Get_WithExpand(t *testing.T) {
	expectResponse("phoneNumberComplianceGetResponse.json", 200)
	complianceId := "767bc62c-2332-4a34-959c-1f6416186254"

	resp, err := client.PhoneNumberCompliance.Get(complianceId, PhoneNumberComplianceGetParams{
		Expand: "end_user,documents,linked_numbers",
	})
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assert.Equal(t, complianceId, resp.Compliance.ComplianceID)

	assertRequest(t, "GET", "PhoneNumber/Compliance/%s", complianceId)
}

func TestPhoneNumberComplianceService_Get_NilClient(t *testing.T) {
	expectResponse("phoneNumberComplianceGetResponse.json", 200)

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumberCompliance.Get("some-id", PhoneNumberComplianceGetParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

// ── Compliance List ──────────────────────────────────────────────────────────

func TestPhoneNumberComplianceService_List(t *testing.T) {
	expectResponse("phoneNumberComplianceListResponse.json", 200)

	resp, err := client.PhoneNumberCompliance.List(PhoneNumberComplianceListParams{})
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assert.Equal(t, 2, len(resp.Compliances))
	assert.Equal(t, 20, resp.Meta.Limit)
	assert.Equal(t, 0, resp.Meta.Offset)
	assert.Equal(t, 2, resp.Meta.TotalCount)
	assert.Equal(t, "074d687f-5630-483d-8349-5b9d7686d673", resp.Compliances[0].ComplianceID)
	assert.Equal(t, "accepted", resp.Compliances[0].Status)
	assert.Equal(t, "767bc62c-2332-4a34-959c-1f6416186254", resp.Compliances[1].ComplianceID)
	assert.Equal(t, "submitted", resp.Compliances[1].Status)

	assertRequest(t, "GET", "PhoneNumber/Compliance")
}

func TestPhoneNumberComplianceService_List_WithFilters(t *testing.T) {
	expectResponse("phoneNumberComplianceListResponse.json", 200)

	resp, err := client.PhoneNumberCompliance.List(PhoneNumberComplianceListParams{
		Limit:      5,
		Offset:     0,
		Status:     "submitted",
		CountryISO: "IN",
		NumberType: "local",
		UserType:   "business",
		Alias:      "TestBiz",
		Expand:     "end_user",
	})
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assertRequest(t, "GET", "PhoneNumber/Compliance")
}

func TestPhoneNumberComplianceService_List_NilClient(t *testing.T) {
	expectResponse("phoneNumberComplianceListResponse.json", 200)

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumberCompliance.List(PhoneNumberComplianceListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

// ── Compliance Delete ────────────────────────────────────────────────────────

func TestPhoneNumberComplianceService_Delete(t *testing.T) {
	expectResponse("phoneNumberComplianceDeleteResponse.json", 200)
	complianceId := "d73b0188-08a8-4bb0-8acf-25e23e274120"

	resp, err := client.PhoneNumberCompliance.Delete(complianceId)
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assert.Equal(t, complianceId, resp.ComplianceID)
	assert.Equal(t, "Compliance application deleted.", resp.Message)

	assertRequest(t, "DELETE", "PhoneNumber/Compliance/%s", complianceId)
}

func TestPhoneNumberComplianceService_Delete_NilClient(t *testing.T) {
	expectResponse("phoneNumberComplianceDeleteResponse.json", 200)

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumberCompliance.Delete("some-id")
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

// ── Compliance Link ──────────────────────────────────────────────────────────

func TestPhoneNumberComplianceLinkService_BulkLink(t *testing.T) {
	expectResponse("phoneNumberComplianceLinkResponse.json", 200)

	resp, err := client.PhoneNumberComplianceLink.BulkLink(PhoneNumberComplianceBulkLinkParams{
		Numbers: []PhoneNumberComplianceLinkNumber{
			{Number: "+912248885512", ComplianceApplicationID: "074d687f-5630-483d-8349-5b9d7686d673"},
			{Number: "+912248885513", ComplianceApplicationID: "074d687f-5630-483d-8349-5b9d7686d673"},
		},
	})
	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, resp.ApiID)
	assert.Equal(t, 2, resp.TotalCount)
	assert.Equal(t, 1, resp.UpdatedCount)
	assert.Equal(t, 2, len(resp.Report))
	assert.Equal(t, "+912248885512", resp.Report[0].Number)
	assert.Equal(t, "success", resp.Report[0].Status)
	assert.Equal(t, "+912248885513", resp.Report[1].Number)
	assert.Equal(t, "failed", resp.Report[1].Status)

	assertRequest(t, "POST", "PhoneNumber/Compliance/Link")
}

func TestPhoneNumberComplianceLinkService_BulkLink_NilClient(t *testing.T) {
	expectResponse("phoneNumberComplianceLinkResponse.json", 200)

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumberComplianceLink.BulkLink(PhoneNumberComplianceBulkLinkParams{
		Numbers: []PhoneNumberComplianceLinkNumber{
			{Number: "+912248885512", ComplianceApplicationID: "some-id"},
		},
	})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

// ── Edge Cases ───────────────────────────────────────────────────────────────

func TestPhoneNumberComplianceService_Get_ErrorResponse(t *testing.T) {
	expectedStatusCode = 404
	expectedResponse = `{"error": "Compliance application not found."}`

	_, err := client.PhoneNumberCompliance.Get("nonexistent-id", PhoneNumberComplianceGetParams{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestPhoneNumberComplianceService_Delete_ErrorResponse(t *testing.T) {
	expectedStatusCode = 404
	expectedResponse = `{"error": "Compliance application not found."}`

	_, err := client.PhoneNumberCompliance.Delete("nonexistent-id")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestPhoneNumberComplianceService_List_EmptyResponse(t *testing.T) {
	expectedStatusCode = 200
	expectedResponse = `{"api_id": "test-id", "meta": {"limit": 20, "offset": 0, "total_count": 0, "next": null, "previous": null}, "compliances": []}`

	resp, err := client.PhoneNumberCompliance.List(PhoneNumberComplianceListParams{})
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 0, len(resp.Compliances))
	assert.Equal(t, 0, resp.Meta.TotalCount)
}

func TestPhoneNumberComplianceLinkService_BulkLink_EmptyReport(t *testing.T) {
	expectedStatusCode = 200
	expectedResponse = `{"api_id": "test-id", "total_count": 0, "updated_count": 0, "report": []}`

	resp, err := client.PhoneNumberComplianceLink.BulkLink(PhoneNumberComplianceBulkLinkParams{
		Numbers: []PhoneNumberComplianceLinkNumber{},
	})
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 0, resp.TotalCount)
	assert.Equal(t, 0, resp.UpdatedCount)
	assert.Equal(t, 0, len(resp.Report))
}

func TestPhoneNumberComplianceRequirementService_Get_EmptyDocTypes(t *testing.T) {
	expectedStatusCode = 200
	expectedResponse = `{"api_id": "test-id", "requirement_id": "req-123", "country_iso": "US", "number_type": "tollfree", "user_type": "business", "document_types": []}`

	resp, err := client.PhoneNumberComplianceRequirements.Get(PhoneNumberComplianceRequirementGetParams{
		CountryISO: "US",
		NumberType: "tollfree",
		UserType:   "business",
	})
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "US", resp.CountryISO)
	assert.Equal(t, 0, len(resp.DocumentTypes))
}

func TestPhoneNumberComplianceService_Get_WithOptionalFields(t *testing.T) {
	expectedStatusCode = 200
	expectedResponse = `{
		"api_id": "test-id",
		"compliance": {
			"compliance_id": "test-comp-id",
			"alias": "test",
			"status": "rejected",
			"country_iso": "IN",
			"number_type": "local",
			"user_type": "business",
			"rejection_reason": "Invalid documents",
			"callback_url": "https://example.com",
			"callback_method": "GET",
			"created_at": "2026-04-06T10:44:17Z",
			"updated_at": "2026-04-06T10:44:17Z",
			"end_user": {
				"end_user_id": "eu-123",
				"type": "business",
				"name": "TestCorp"
			},
			"documents": [
				{
					"document_id": "doc-123",
					"document_type_id": "dt-123",
					"document_name": "Registration Certificate"
				}
			],
			"linked_numbers": [
				{"number": "+911234567890", "number_type": "local"}
			]
		}
	}`

	resp, err := client.PhoneNumberCompliance.Get("test-comp-id", PhoneNumberComplianceGetParams{
		Expand: "end_user,documents,linked_numbers",
	})
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "rejected", resp.Compliance.Status)
	assert.Equal(t, "Invalid documents", resp.Compliance.RejectionReason)
	assert.NotNil(t, resp.Compliance.EndUser)
	assert.Equal(t, "eu-123", resp.Compliance.EndUser.EndUserID)
	assert.Equal(t, 1, len(resp.Compliance.Documents))
	assert.Equal(t, "doc-123", resp.Compliance.Documents[0].DocumentID)
	assert.Equal(t, 1, len(resp.Compliance.LinkedNumbers))
	assert.Equal(t, "+911234567890", resp.Compliance.LinkedNumbers[0].Number)
}
