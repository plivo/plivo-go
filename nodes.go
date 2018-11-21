package plivo

type NodeActionResponse struct {
	ApiID     string `json:"api_id" url:"api_id"`
	Error     string `json:"error" url:"error"`
}
type MultiPartyCallActionPayload struct {
	Action     string `json:"action" url:"action"`
	To     string `json:"to" url:"to"`
	Role     string `json:"role" url:"role"`
	TriggerSource     string `json:"trigger_source" url:"trigger_source"`
}

type Node struct {
	NodeID     string `json:"node_id" url:"node_id"`
	PhloID string `json:"phlo_id" url:"phlo_id"`
	Name string `json:"name" url:"name"`
	NodeType string `json:"node_type" url:"node_type"`
	CreatedOn string `json:"created_on" url:"created_on"`
}

type MultiPartyCall struct {
	Node
	BaseResource
}
type MultiPartyCallMemberActionPayload struct {
	Action string `json:"action" url:"action"`
}
type MultiPartyCallMember struct {
	NodeID     string `json:"node_id" url:"node_id"`
	PhloID string `json:"phlo_id" url:"phlo_id"`
	NodeType string `json:"node_type" url:"node_type"`
	MemberAddress string `json:"member_address" url:"member_address"`
	BaseResource
}

func (self *MultiPartyCall) update(params MultiPartyCallActionPayload) (response *NodeActionResponse, err error) {
	req, err := self.client.NewRequest("POST", params,"phlo/%s/%s/%s", self.PhloID, self.NodeType,
		self.NodeID)
	if err != nil {
		return
	}
	response = &NodeActionResponse{}
	err = self.client.ExecuteRequest(req, response)

	return
}

func (self *MultiPartyCall) Call(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return self.update(params)
}

func (self *MultiPartyCall) WarmTransfer(params MultiPartyCallActionPayload) (response *NodeActionResponse,
	err error) {
	return self.update(params)
}

func (self *MultiPartyCall) ColdTransfer(params MultiPartyCallActionPayload) (response *NodeActionResponse,
	err error) {
	return self.update(params)
}

func (self *MultiPartyCall) Member(memberID string ) (response *MultiPartyCallMember,
	err error) {
	response = &MultiPartyCallMember{self.NodeID, self.PhloID, self. NodeType,memberID, BaseResource{self.client}}
	return
}

func (service *MultiPartyCallMember) AbortTransfer(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return service.update(MultiPartyCallMemberActionPayload{"abort_tranfer"})
}
func (service *MultiPartyCallMember) ResumeCall(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return service.update(MultiPartyCallMemberActionPayload{"resume_call"})
}
func (service *MultiPartyCallMember) VoiceMailDrop(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return service.update(MultiPartyCallMemberActionPayload{"voicemail_drop"})
}
func (service *MultiPartyCallMember) HangUp(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return service.update(MultiPartyCallMemberActionPayload{"hangup"})
}
func (service *MultiPartyCallMember) Hold(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return service.update(MultiPartyCallMemberActionPayload{"hold"})
}
func (service *MultiPartyCallMember) UnHold(params MultiPartyCallActionPayload) (*NodeActionResponse, error) {
	return service.update(MultiPartyCallMemberActionPayload{"unhold"})
}

func (service *MultiPartyCallMember) update(params interface{}) (response *NodeActionResponse, err error) {
	req, err := service.client.NewRequest("POST", params,"phlo/%s/%s/%s", service.PhloID, service.NodeType,
		service.NodeID)
	if err != nil {
		return
	}
	response = &NodeActionResponse{}
	err = service.client.ExecuteRequest(req, response)

	return
}