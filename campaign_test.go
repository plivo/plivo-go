package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCampaign_List(t *testing.T) {
	expectResponse("campaignListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Campaign.List(CampaignListParams{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.CampaignResponse[0].CampaignID)
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
	CampaignID := "CMPT4EP"
	assert := require.New(t)
	campaign, err := client.Campaign.Get(CampaignID)
	assert.NotNil(campaign)
	assert.Nil(err)
	assert.Equal(CampaignID, campaign.Campaign.CampaignID)
	assert.NotEmpty(campaign.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	campaign, err = client.Campaign.Get(CampaignID)
	assert.NotNil(err)
	assert.Nil(campaign)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Campaign/%s", CampaignID)
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
	resp, err := client.Campaign.Create(CampaignCreationParams{
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
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.Campaign.CampaignID)
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
	assert.NotEmpty(resp.Campaign.CampaignID)
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
