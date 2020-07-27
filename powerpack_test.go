package plivo

import (
	"errors"
	"testing"
)

func TestPowerpack_List(t *testing.T) {
	expectResponse("powerpackList.json", 202)

	if _, err := client.Powerpack.List(PowerpackListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.List(PowerpackListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Powerpack/")
}

func TestPowerpack_Get(t *testing.T) {
	expectResponse("powerpack.json", 200)
	powerpackUUID := "powerpackUUID"

	if _, err := client.Powerpack.Get(powerpackUUID); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Get(powerpackUUID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Powerpack/%s/", powerpackUUID)
}

func TestPowerpack_Create(t *testing.T) {
	expectResponse("powerpack.json", 200)

	if _, err := client.Powerpack.Create(PowerackCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Create(PowerackCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Powerpack/")
}

func TestPowerpack_Update(t *testing.T) {
	expectResponse("powerpack.json", 200)
	powerpackuuid := "powerpackuuid"
	if _, err := client.Powerpack.Update(PowerackUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Update(PowerackUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Powerpack/%s/", powerpackuuid)
}

func TestPowerpack_delete(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	powerpackUUID := "powerpackUUID"
	if _, err := client.Powerpack.Delete(PowerpackDeleteParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Delete(PowerpackDeleteParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Powerpack/%s/", powerpackUUID)
}

func TestPowerpack_find_number(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	number := "number"
	if _, err := client.Powerpack.Find_numbers(number); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Find_numbers(number)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/%s/", numberpool_uuid, number)
}

func TestPowerpack_FindNumbersWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	params := PowerpackFindNumberOptions{
		Service: MMS,
	}
	number := "number"
	if _, err := client.Powerpack.FindNumbersWithOptions(number, params); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.FindNumbersWithOptions(number, params)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/%s/", numberpool_uuid, number)
}

func TestPowerpack_AddNumber(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	number := "number"
	if _, err := client.Powerpack.Add_number(number); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Add_number(number)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "NumberPool/%s/%s/", numberpool_uuid, number)
}

func TestPowerpack_AddNumberWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	params := PowerpackAddNumberOptions{
		Service: MMS,
	}
	number := "number"
	if _, err := client.Powerpack.AddNumberWithOptions(number, params); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.AddNumberWithOptions(number, params)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "NumberPool/%s/%s/", numberpool_uuid, number)
}

func TestPowerpack_RemoveNumber(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	number := "number"
	if _, err := client.Powerpack.Remove_number(number, NumberRemoveParams{Unrent: false}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Remove_number(number, NumberRemoveParams{Unrent: false})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "NumberPool/%s/%s/", numberpool_uuid, number)
}

func TestPowerpack_ListNumbers(t *testing.T) {
	expectResponse("numberpoolNumberRespone.json", 200)
	numberpooluuid := "numberpool_uuid"
	if _, err := client.Powerpack.List_numbers(PowerpackSearchParam{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.List_numbers(PowerpackSearchParam{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Number/", numberpooluuid)

}

func TestListShortCode(t *testing.T) {
	expectResponse("numberpoolShortCodeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	if _, err := client.Powerpack.List_shortcodes(); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.List_shortcodes()
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Shortcode/", numberpooluuid)

}

func TestListShortCodeWithOptions(t *testing.T) {
	expectResponse("numberpoolShortCodeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	params := PowerpackListShortcodeOptions{
		Service: MMS,
	}
	if _, err := client.Powerpack.ListShortcodesWithOptions(params); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.ListShortcodesWithOptions(params)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Shortcode/", numberpooluuid)

}

func TestFindShortCode(t *testing.T) {
	expectResponse("numberpoolSingleShortcodeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	shortcode := "shortcode"
	if _, err := client.Powerpack.Find_shortcode(shortcode); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Find_shortcode(shortcode)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Shortcode/%s/", numberpooluuid, shortcode)

}

func TestFindShortCodeWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleShortcodeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	param := PowerpackFindShortcodeOptions{
		Service: MMS,
	}
	shortcode := "shortcode"
	if _, err := client.Powerpack.FindShortcodeWithOptions(shortcode, param); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.FindShortcodeWithOptions(shortcode, param)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Shortcode/%s/", numberpooluuid, shortcode)

}
func TestBuyAddNumber(t *testing.T) {
	expectResponse("numberpoolSingleNoResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	number := "number"
	if _, err := client.Powerpack.Buy_add_number(BuyPhoneNumberParam{Number: "15795890617"}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Buy_add_number(BuyPhoneNumberParam{Number: "15795890617"})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
	assertRequest(t, "GET", "NumberPool/%s/Number/%s/", numberpooluuid, number)
}

func TestFindTollfree(t *testing.T) {
	expectResponse("numberpoolSingleTollfreeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	tollfree := "tollfree"
	if _, err := client.Powerpack.Find_tollfree(tollfree); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Find_tollfree(tollfree)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Tollfree/%s/", numberpooluuid, tollfree)

}

func TestFindTollfreeWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleTollfreeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	params := PowerpackFindTollfreeOptions{
		Service: MMS,
	}
	tollfree := "tollfree"
	if _, err := client.Powerpack.FindTollfreeWithOptions(tollfree, params); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.FindTollfreeWithOptions(tollfree, params)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Tollfree/%s/", numberpooluuid, tollfree)

}

func TestListTollfree(t *testing.T) {
	expectResponse("numberpoolTollfreeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	if _, err := client.Powerpack.List_tollfree(); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.List_tollfree()
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Tollfree/", numberpooluuid)

}

func TestListTollfreeWithOptions(t *testing.T) {
	expectResponse("numberpoolTollfreeResponse.json", 200)
	numberpooluuid := "numberpool_uuid"
	params := PowerpackListTollfreeOptions{
		Service: MMS,
	}
	if _, err := client.Powerpack.ListTollfreeWithOptions(params); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.ListTollfreeWithOptions(params)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "NumberPool/%s/Tollfree/", numberpooluuid)

}

func TestPowerpack_RemoveTollfree(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	tollfree := "tollfree"
	if _, err := client.Powerpack.Remove_tollfree(tollfree, true); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Remove_tollfree(tollfree, true)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "NumberPool/%s/Tollfree/%s/", numberpool_uuid, tollfree)
}

func TestPowerpack_RemoveShortcode(t *testing.T) {
	expectResponse("powerpackDeleteResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	shortcode := "shortcode"
	if _, err := client.Powerpack.Remove_shortcode(shortcode); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Remove_shortcode(shortcode, true)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "NumberPool/%s/Shortcode/%s/", numberpool_uuid, number)
}

func TestPowerpack_AddTollfree(t *testing.T) {
	expectResponse("numberpoolSingleTollfreeResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	tollfree := "tollfree"
	if _, err := client.Powerpack.Add_tollfree(tollfree); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.Add_tollfree(tollfree)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "NumberPool/%s/Tollfree/%s/", numberpool_uuid, tollfree)
}

func TestPowerpack_AddTollfreeWithOptions(t *testing.T) {
	expectResponse("numberpoolSingleTollfreeResponse.json", 200)
	numberpool_uuid := "numberpool_uuid"
	params := PowerpackAddTollfreeOptions{
		Service: MMS,
	}
	tollfree := "tollfree"
	if _, err := client.Powerpack.AddTollfreeWithOptions(tollfree, params); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Powerpack.AddTollfreeWithOptions(tollfree, params)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "NumberPool/%s/Tollfree/%s/", numberpool_uuid, tollfree)
}
