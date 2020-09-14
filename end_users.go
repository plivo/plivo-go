package plivo

import "time"

type EndUserService struct {
	client *Client
}

type EndUserGetResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	EndUserID   string    `json:"end_user_id"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	EndUserType string    `json:"end_user_type"`
	APIID       string    `json:"api_id"`
}

func (service *EndUserService) Get(endUserId string) (response *EndUserGetResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "EndUser/%s", endUserId)
	response = &EndUserGetResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}
