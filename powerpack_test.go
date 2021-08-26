package plivo

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPowerpack_List(t *testing.T) {
	expectResponse("powerpackList.json", 202)
	assert := require.New(t)
	resp, err := client.Powerpack.List(PowerpackListParams{})
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Meta.TotalCount)
	assert.NotNil(resp.Objects[0].Name)
	assert.NotNil(resp.Objects[0].UUID)
	assertBaseRequest(t, http.MethodGet, "/v1/Account/AuthId/Powerpack//")

	// with httpclient nill
	cl := client.httpClient
	client.httpClient = nil
	res, err := client.Powerpack.List(PowerpackListParams{})
	assert.NotNil(err)
	assert.Nil(res)
	client.httpClient = cl
	assert.Equal("client and httpClient cannot be nil", err.Error())
	assertRequest(t, "GET", "Powerpack/")
}

func TestPowerpack_ListWithParam(t *testing.T) {
	expectResponse("powerpackList.json", 202)
	assert := require.New(t)
	response, err := client.Powerpack.List(PowerpackListParams{Limit: 10, Offset: 0, Service: "sms"})
	assert.Nil(err)
	assert.NotNil(response)
	assertRequest(t, "GET", "Powerpack/")
}

func TestPowerpack_Get(t *testing.T) {
	expectResponse("powerpack.json", 200)
	assert := require.New(t)
	powerpackUUID := "powerpackUUID"

	resp, err := client.Powerpack.Get(powerpackUUID)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Name)
	assert.NotNil(resp.NumberPoolUUID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Get(powerpackUUID)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl
	assertRequest(t, "GET", "Powerpack/%s/", powerpackUUID)
}

func TestPowerpack_Create(t *testing.T) {
	expectResponse("powerpack.json", 200)
	assert := require.New(t)
	resp, err := client.Powerpack.Create(PowerackCreateParams{})
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Name)
	assert.NotNil(resp.ApiID)
	assert.NotNil(resp.UUID)
	response, err := client.Powerpack.Create(PowerackCreateParams{Name: "vishnu_sep_01"})
	assert.Nil(err)
	assert.Equal(response.Name, "vishnu_sep_01")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Create(PowerackCreateParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl
	assertRequest(t, "POST", "Powerpack")
}

func TestPowerpack_Update(t *testing.T) {
	expectResponse("powerpack.json", 200)
	assert := require.New(t)
	powerpackUUID := "86bbb125-97bb-4d72-89fd-81d5c515b015"
	client.Powerpack.UUID = powerpackUUID
	response, err := client.Powerpack.Update(PowerackUpdateParams{})
	assert.Nil(err)

	assert.Equal(response.StickySender, true)
	assert.Equal(response.Name, "vishnu_sep_01")
	cl := client.httpClient
	client.httpClient = nil
	response, err = client.Powerpack.Update(PowerackUpdateParams{})
	assert.NotNil(err)
	assert.Nil(response)
	client.httpClient = cl
	assertRequest(t, "POST", "Powerpack/%s", powerpackUUID)
}

func TestPowerpack_Delete(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	assert := require.New(t)
	powerpackUUID := "86bbb125-97bb-4d72-89fd-81d5c515b015"
	client.Powerpack.UUID = powerpackUUID
	res, err := client.Powerpack.Delete(PowerpackDeleteParams{})
	assert.NotNil(res)
	assert.Nil(err)
	assert.Equal(res.Response, "success")
	assert.NotEmpty(res.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	res, err = client.Powerpack.Delete(PowerpackDeleteParams{})
	assert.NotNil(err)
	assert.Nil(res)
	client.httpClient = cl
	assertRequest(t, "DELETE", "Powerpack/%s", powerpackUUID)
}

func TestPowerpack_Find_numbers(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	assert := require.New(t)
	number := "15799140348"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	resp, err := client.Powerpack.Find_numbers(number)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Number, number)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Find_numbers(number)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl
}

func TestPowerpack_FindNumbersWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	params := PowerpackFindNumberOptions{
		Service: MMS,
	}
	number := "15799140348"
	resp, err := client.Powerpack.FindNumbersWithOptions(number, params)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Service, "mms")

	resp, err = client.Powerpack.FindNumbersWithOptions(number, params)
	assert.Nil(err)
	assert.NotNil(resp)

	assertBaseRequest(t, "GET", "/v1/Account/AuthId/NumberPool/%s/Number/%s//?service=mms", numberpool_uuid, number)
}

func TestPowerpack_Add_number(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	number := "15799140348"
	resp, err := client.Powerpack.Add_number(number)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Number, number)
	assert.Equal(resp.Type, "fixed")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Add_number(number)
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "POST", "NumberPool/%s/Number/%s", numberpool_uuid, number)
}

func TestPowerpack_Add_numberWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	params := PowerpackAddNumberOptions{
		Service: MMS,
	}
	number := "15799140348"
	resp, err := client.Powerpack.AddNumberWithOptions(number, params)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Service, "mms")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.AddNumberWithOptions(number, params)
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "POST", "NumberPool/%s/Number/%s", numberpool_uuid, number)
}

func TestPowerpack_Remove_number(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	number := "15799140348"
	resp, err := client.Powerpack.Remove_number(number, NumberRemoveParams{Unrent: false})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Response, "success")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Remove_number(number, NumberRemoveParams{Unrent: false})
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "DELETE", "NumberPool/%s/Number/%s", numberpool_uuid, number)
}

func TestPowerpack_List_numbers(t *testing.T) {
	expectResponse("numberpoolNumberRespone.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	resp, err := client.Powerpack.List_numbers(PowerpackSearchParam{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Objects[0].Number_pool_uuid, numberpool_uuid)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.List_numbers(PowerpackSearchParam{})
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "GET", "NumberPool/%s/Number/", numberpool_uuid)
}

func TestList_shortcodes(t *testing.T) {
	expectResponse("numberpoolShortCodeResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	resp, err := client.Powerpack.List_shortcodes()
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Objects[0].Number_pool_uuid, numberpool_uuid)
	assert.NotEmpty(resp.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.List_shortcodes()
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "GET", "NumberPool/%s/Shortcode", numberpool_uuid)

}

func TestFind_shortcode(t *testing.T) {
	expectResponse("numberpoolSingleShortcodeResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	shortcode := "444444"
	resp, err := client.Powerpack.Find_shortcode(shortcode)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.ShortCode.Shortcode, shortcode)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Find_shortcode(shortcode)
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "GET", "NumberPool/%s/Shortcode/%s/", numberpool_uuid, shortcode)

}

func TestBuy_add_number(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	number := "15799140348"
	resp, err := client.Powerpack.Buy_add_number(BuyPhoneNumberParam{Number: number})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Number, number)

	assertRequest(t, "POST", "NumberPool/%s/Number/%s", numberpool_uuid, number)
}

func TestFind_tollfree(t *testing.T) {
	expectResponse("numberpoolSingleTollfreeResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	tollfree := "18772209942"
	resp, err := client.Powerpack.Find_tollfree(tollfree)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Tollfree.Tollfree, tollfree)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Find_tollfree(tollfree)
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "GET", "NumberPool/%s/Tollfree/%s/", numberpool_uuid, tollfree)

}

func TestList_tollfree(t *testing.T) {
	expectResponse("numberpoolTollfreeResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	resp, err := client.Powerpack.List_tollfree()
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.Objects[0].Tollfree)
	assert.NotNil(resp.Objects)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.List_tollfree()
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "GET", "NumberPool/%s/Tollfree", numberpool_uuid)

}

func TestRemove_tollfree(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	tollfree := "18772209942"
	resp, err := client.Powerpack.Remove_tollfree(tollfree, NumberRemoveParams{Unrent: false})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Response, "success")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Remove_tollfree(tollfree, NumberRemoveParams{Unrent: false})
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "DELETE", "NumberPool/%s/Tollfree/%s", numberpool_uuid, tollfree)
}

func TestRemove_shortcode(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	shortcode := "444444"
	resp, err := client.Powerpack.Remove_shortcode(shortcode)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Response, "success")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Remove_shortcode(shortcode)
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "DELETE", "NumberPool/%s/Shortcode/%s", numberpool_uuid, shortcode)
}

func TestAdd_tollfree(t *testing.T) {
	expectResponse("numberpoolSingleTollfreeResponse.json", 200)
	assert := require.New(t)
	numberpool_uuid := "659c7f88-c819-46e2-8af4-2d8a84249099"
	client.Powerpack.NumberPoolUUID = "/v1/Account/MAODZKMDFJMJU3MTEYNG/NumberPool/659c7f88-c819-46e2-8af4-2d8a84249099/"
	tollfree := "18772209942"
	resp, err := client.Powerpack.Add_tollfree(tollfree)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.Number, tollfree)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Powerpack.Add_tollfree(tollfree)
	assert.Nil(resp)
	assert.NotNil(err)
	client.httpClient = cl
	assertRequest(t, "POST", "NumberPool/%s/Tollfree/%s", numberpool_uuid, tollfree)
}
