package plivo

type Phlo struct {
	BaseResource

	PhloId     string `json:"phlo_id" url:"phlo_id"`
	Name     string `json:"name" url:"name"`
	CreatedOn string `json:"created_on" url:"created_on"`
}

type Phlos struct {
	BaseResourceInterface
}

func NewPhlos(client *PhloClient) (phlos *Phlos){
	phlos = &Phlos{}
	phlos.client = client
	//phlos.resourceType = Phlo // TODO: how to set this?

	return
}

type PhloRun struct {
	ApiID     string `json:"api_id" url:"api_id"`
	PhloRunID string `json:"phlo_run_id" url:"phlo_run_id"`
	PhloID string `json:"phlo_id" url:"phlo_id"`
	Message string `json:"message" url:"message"`
}


func (self *Phlos) Get(phloId string) (response *Phlo, err error) {
	req, err := self.client.NewRequest("GET", nil,"phlo/%s", phloId)
	response = &Phlo{}
	response.client = self.client // Todo: set this in ExecuteRequest()
	err = self.client.ExecuteRequest(req, response)

	return
}


func (self *Phlo) Node(phloId string, nodeType string, nodeId string) (response *Node, err error) {
	req, err := self.client.NewRequest("GET", nil,"phlo/%s/%s/%s", phloId, nodeType, nodeId)
	response = &Node{}
	err = self.client.ExecuteRequest(req, response)

	return
}

func (self *Phlo) MultiPartyCall(phloId string, nodeType string, nodeId string) (response *MultiPartyCall, err error) {
	req, err := self.client.NewRequest("GET", nil,"phlo/%s/%s/%s", phloId, nodeType, nodeId)
	response = &MultiPartyCall{}
	response.client = self.client
	err = self.client.ExecuteRequest(req, response)
	return
}

func (self *Phlo) Run(data map[string]interface{}) (response *PhloRun, err error) {
	req, err := self.client.NewRequest("POST", data,"account/%s/phlo/%s/", self.client.AuthId, self.PhloId)
	response = &PhloRun{}
	err = self.client.ExecuteRequest(req, response)
	return
}