package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCampaign_List(t *testing.T) {
	expectResponse("campaignListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Campaign.List(CampaignListParams{Limit: 2, Offset: 0})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.CampaignResponse[0].CampaignID)
	assert.Equal(false, resp.CampaignResponse[0].CampaignAttributes.SubscriberOptin)
	cl := client.httpClient
	client.httpClient = nil
	brandID := ""
	usercase := "MIXED"
	resp, err = client.Campaign.List(CampaignListParams{BrandID: &brandID, Usecase: &usercase})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Campaign")
}

func TestCampaign_Get(t *testing.T) {
	expectResponse("campaignGetResponse.json", 200)
	CampaignID := "CY5NVUA"
	SubscriberOptin := false
	assert := require.New(t)
	campaign, err := client.Campaign.Get(CampaignID)
	assert.NotNil(campaign)
	assert.Nil(err)
	assert.Equal(CampaignID, campaign.Campaign.CampaignID)
	assert.Equal(SubscriberOptin, campaign.Campaign.CampaignAttributes.SubscriberOptin)
	assert.NotEmpty(campaign.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	campaign, err = client.Campaign.Get(CampaignID)
	assert.NotNil(err)
	assert.Nil(campaign)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Campaign/%s", CampaignID)
}

func TestImportCampaign(t *testing.T) {
	expectResponse("campaignImportResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Campaign.Import(ImportCampaignParams{
		CampaignID:    "C1QGYD1",
		CampaignAlias: "import campaign",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.CampaignID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Campaign.Import(ImportCampaignParams{
		CampaignID:    "C1QGYD1",
		CampaignAlias: "import campaign",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "10dlc/Campaign/Import")
}

func TestCampaign_Create(t *testing.T) {
	expectResponse("campaignCreationResponse.json", 200)
	assert := require.New(t)
	campaignAlias := "campaign name sssample"
	embeddedLink := false
	embeddedPhone := false
	ageGated := false
	directLending := false
	subUsecases := []string{"CUSTOMER_CARE", "2FA"}
	sample1 := "test1"
	sample2 := "test2"
	resp, err := client.Campaign.Create(CampaignCreationParams{
		BrandID:            "B8OD95Z",
		CampaignAlias:      &campaignAlias,
		Vertical:           "INSURANCE",
		Usecase:            "MIXED",
		SubUsecases:        &subUsecases,
		Description:        "sample description text",
		EmbeddedLink:       &embeddedLink,
		EmbeddedPhone:      &embeddedPhone,
		AgeGated:           &ageGated,
		DirectLending:      &directLending,
		SubscriberOptin:    true,
		SubscriberOptout:   true,
		SubscriberHelp:     true,
		Sample1:            &sample1,
		Sample2:            &sample2,
		AffiliateMarketing: false,
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.CampaignID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Campaign.Create(CampaignCreationParams{
		BrandID:            "B8OD95Z",
		CampaignAlias:      &campaignAlias,
		Vertical:           "INSURANCE",
		Usecase:            "MIXED",
		SubUsecases:        &subUsecases,
		Description:        "sample description text",
		EmbeddedLink:       &embeddedLink,
		EmbeddedPhone:      &embeddedPhone,
		AgeGated:           &ageGated,
		DirectLending:      &directLending,
		SubscriberOptin:    true,
		SubscriberOptout:   true,
		SubscriberHelp:     true,
		Sample1:            &sample1,
		AffiliateMarketing: false,
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "10dlc/Campaign")
}

func TestCampaign_Create_WithMoreAttribute(t *testing.T) {
	expectResponse("campaignCreationResponse.json", 200)
	assert := require.New(t)
	campaignAlias := "campaign name sssample"
	embeddedLink := false
	embeddedPhone := false
	ageGated := false
	directLending := false
	subUsecases := []string{"CUSTOMER_CARE", "2FA"}
	sample1 := "sample 1 should be minimum 20 character"
	sample2 := "sample 2 should be minimum 20 character"

	resp, err := client.Campaign.Create(CampaignCreationParams{
		BrandID:          "B8OD95Z",
		CampaignAlias:    &campaignAlias,
		Vertical:         "INSURANCE",
		Usecase:          "MIXED",
		SubUsecases:      &subUsecases,
		Description:      "sample description text should be minimum 40 character, should be minimum 40 character,should be minimum 40 character",
		EmbeddedLink:     &embeddedLink,
		EmbeddedPhone:    &embeddedPhone,
		AgeGated:         &ageGated,
		DirectLending:    &directLending,
		SubscriberOptin:  true,
		SubscriberOptout: true,
		SubscriberHelp:   true,
		Sample1:          &sample1,
		Sample2:          &sample2,
		MessageFlow:      "message_flow should be minimum 40 character,message_flow should be minimum 40 character",
		OptoutMessage:    "optout_message should be mandatory param",
		HelpMessage:      "help_message should be mandatory param",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Campaign.Create(CampaignCreationParams{
		BrandID:          "B8OD95Z",
		CampaignAlias:    &campaignAlias,
		Vertical:         "INSURANCE",
		Usecase:          "MIXED",
		SubUsecases:      &subUsecases,
		Description:      "sample description text",
		EmbeddedLink:     &embeddedLink,
		EmbeddedPhone:    &embeddedPhone,
		AgeGated:         &ageGated,
		DirectLending:    &directLending,
		SubscriberOptin:  true,
		SubscriberOptout: true,
		SubscriberHelp:   true,
		Sample1:          &sample1,
		Sample2:          &sample2,
		MessageFlow:      "message_flow should be minimum 40 character,message_flow should be minimum 40 character",
		OptoutMessage:    "optout_message should be mandatory param",
		HelpMessage:      "help_message should be mandatory param",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "10dlc/Campaign")
}

func TestCampaign_Update(t *testing.T) {
	expectResponse("campaignUpdateResponse.json", 200)
	campaignID := "CXNSG9W"
	assert := require.New(t)
	sample1 := "test1"
	sample2 := "test2"
	resp, err := client.Campaign.Update(campaignID, CampaignUpdateParams{
		Sample1: sample1,
		Sample2: sample2,
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.Campaign.CampaignID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Campaign.Update(campaignID, CampaignUpdateParams{
		Sample1: sample1,
		Sample2: sample2,
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "10dlc/Campaign/%s", campaignID)
}

func TestCampaign_NumberLink(t *testing.T) {
	expectResponse("campaignNumberLinkUnlinkResponse.json", 200)
	assert := require.New(t)
	numbers := []string{"9876564598", "78654567876"}
	resp, err := client.Campaign.NumberLink("CY5NVUA", CampaignNumberLinkParams{
		Numbers: numbers,
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Campaign.NumberLink("CY5NVUA", CampaignNumberLinkParams{
		Numbers: numbers,
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "10dlc/Campaign/CY5NVUA/Number")
}

func TestCampaign_NumberGet(t *testing.T) {
	expectResponse("campaignNumberGetResponse.json", 200)
	CampaignID := "CY5NVUA"
	number := "14845007032"
	assert := require.New(t)
	campaign, err := client.Campaign.NumberGet(CampaignID, number)
	assert.NotNil(campaign)
	assert.Nil(err)
	assert.Equal(CampaignID, campaign.CampaignID)
	assert.Equal(number, campaign.CampaignNumbers[0].Number)
	assert.NotEmpty(campaign.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	campaign, err = client.Campaign.NumberGet(CampaignID, number)
	assert.NotNil(err)
	assert.Nil(campaign)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Campaign/%s/Number/%s", CampaignID, number)
}

func TestCampaign_NumbersGet(t *testing.T) {
	expectResponse("campaignNumbersGetResponse.json", 200)
	CampaignID := "CY5NVUA"
	number := "14845007032"
	numberCompletedStatus := 1
	assert := require.New(t)
	campaign, err := client.Campaign.NumbersGet(CampaignID, CampaignNumbersGetParams{Limit: 20, Offset: 0})
	assert.NotNil(campaign)
	assert.Nil(err)
	assert.Equal(CampaignID, campaign.CampaignID)
	assert.Equal(number, campaign.CampaignNumbers[0].Number)
	assert.Equal(numberCompletedStatus, campaign.CampaignNumberSummary["COMPLETED"])
	assert.NotEmpty(campaign.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	campaign, err = client.Campaign.NumbersGet(CampaignID, CampaignNumbersGetParams{Limit: 20, Offset: 0})
	assert.NotNil(err)
	assert.Nil(campaign)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Campaign/%s/Number", CampaignID)
}

func TestCampaign_NumberUnlink(t *testing.T) {
	expectResponse("campaignNumberLinkUnlinkResponse.json", 200)
	assert := require.New(t)
	number := "9876564598"
	resp, err := client.Campaign.NumberUnlink("CY5NVUA", number, CampaignNumberUnlinkParams{
		Method: "POST",
		URL:    "http://example.com/test"})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Campaign.NumberUnlink("CY5NVUA", number, CampaignNumberUnlinkParams{
		Method: "POST",
		URL:    "http://example.com/test"})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "DELETE", "10dlc/Campaign/CY5NVUA/Number/9876564598")
}

func TestCampaign_Delete(t *testing.T) {
	expectResponse("campaignDeleteResponse.json", 200)
	CampaignID := "CY5NVUA"
	assert := require.New(t)
	campaign, err := client.Campaign.Delete(CampaignID)
	assert.NotNil(campaign)
	assert.Nil(err)
	assert.Equal(CampaignID, campaign.CampaignID)
	assert.NotEmpty(campaign.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	campaign, err = client.Campaign.Delete(CampaignID)
	assert.NotNil(err)
	assert.Nil(campaign)
	client.httpClient = cl

	assertRequest(t, "DELETE", "10dlc/Campaign/%s", CampaignID)
}
