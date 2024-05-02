package plivo

import (
	"errors"
	"testing"
)

func TestMaskingSessionService_CreateMaskingSession_GeoMatch(t *testing.T) {
	t.Run("GeoMatch::False", func(t *testing.T) {
		falseValue := false
		_, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{
			FirstParty:  "+919999323467",
			SecondParty: "+919891865130",
			GeoMatch:    &falseValue,
		})
		if err != nil {
			t.Logf("err :: %v", err)
		}
	})

	t.Run("GeoMatch::True", func(t *testing.T) {
		trueValue := true
		_, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{
			FirstParty:  "+919999323467",
			SecondParty: "+919891865130",
			GeoMatch:    &trueValue,
		})
		if err != nil {
			t.Logf("err :: %v", err)
		}
	})

	t.Run("GeoMatch::nil", func(t *testing.T) {
		_, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{
			FirstParty:  "+919999323467",
			SecondParty: "+919891865130",
			GeoMatch:    nil,
		})
		if err != nil {
			t.Logf("err :: %v", err)
		}
	})
}

func TestMaskingSessionService_CreateMaskingSession_SubAccount(t *testing.T) {
	t.Run("SubAccount::Empty", func(t *testing.T) {
		_, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{
			FirstParty:  "+919999323467",
			SecondParty: "+919891865130",
		})
		if err != nil {
			t.Logf("err :: %v", err)
		}
	})

	t.Run("SubAccount::Valid SubAccount", func(t *testing.T) {
		_, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{
			FirstParty:  "+919999323467",
			SecondParty: "+919891865130",
			SubAccount:  "SAZTA0ZJJHMDETOWQ4YI",
		})
		if err != nil {
			t.Logf("err :: %v", err)
		}
	})
}

func TestMaskingSessionService_CreateMaskingSession(t *testing.T) {
	expectResponse("createMaskingSessionResponse.json", 200)

	if _, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.MaskingSession.CreateMaskingSession(CreateMaskingSessionParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Masking/Session")
}

func TestMaskingSessionService_DeleteMaskingSession(t *testing.T) {
	expectResponse("deleteMaskingSessionResponse.json", 204)
	SessionUuid := "15e4256c-be01-475c-9a69-95cf65bbed71"

	if _, err := client.MaskingSession.DeleteMaskingSession(SessionUuid); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.MaskingSession.DeleteMaskingSession(SessionUuid)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Masking/Session/%s", SessionUuid)
}

func TestMaskingSessionService_GetMaskingSession(t *testing.T) {
	expectResponse("getMaskingSessionResponse.json", 200)
	SessionUuid := "15e4256c-be01-475c-9a69-95cf65bbed71"

	if _, err := client.MaskingSession.GetMaskingSession(SessionUuid); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.MaskingSession.GetMaskingSession(SessionUuid)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Masking/Session/%s", SessionUuid)
}

func TestMaskingSessionService_updateMaskingSession(t *testing.T) {
	expectResponse("updateMaskingSessionResponse.json", 204)
	SessionUuid := "15e4256c-be01-475c-9a69-95cf65bbed71"

	if _, err := client.MaskingSession.UpdateMaskingSession(UpdateMaskingSessionParams{}, SessionUuid); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.MaskingSession.UpdateMaskingSession(UpdateMaskingSessionParams{}, SessionUuid)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Masking/Session/%s", SessionUuid)
}

func TestMaskingSessionService_ListMaskingSession(t *testing.T) {
	expectResponse("listMaskingSessionResponse.json", 204)

	if _, err := client.MaskingSession.ListMaskingSession(ListSessionFilterParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.MaskingSession.ListMaskingSession(ListSessionFilterParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Masking/Session")
}
