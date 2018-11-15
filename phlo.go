package plivo

type PhloService struct {
	client *Client
}

type Phlo struct {
	PhloId     string `json:"phlo_id" url:"phlo_id"`
	Name     string `json:"name" url:"name"`
	CreatedOn string `json:"created_on" url:"created_on"`
	client *Client
}


func (service *PhloService) Get(phlo_id string) (response *Phlo, err error) {
	req, err := service.client.NewRequestPhlo("GET", nil,"%s", phlo_id)
	response = &Phlo{client:service.client}
	err = service.client.ExecuteRequest(req, response)
	return
}


func (service *Phlo) GetNode(phlo_id string,node_type string,node_id string) (response *Node, err error) {
	req, err := service.client.NewRequestPhlo("GET", nil,"%s/%s/%s", phlo_id,node_type,node_id)
	response = &Node{}
	err = service.client.ExecuteRequest(req, response)
	return
}

