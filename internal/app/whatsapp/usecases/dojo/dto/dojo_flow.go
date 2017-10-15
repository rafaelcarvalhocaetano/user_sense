package dto

type InputCreateDojoFlow struct {
	Name    *string     `json:"name"`
	Default *string     `json:"default"`
	Flows   []*DojoFlow `json:"flows"`
	Type    *string     `json:"type"`
}

type DojoFlow struct {
	Type            *string `json:"type"`
	Current         *string `json:"current"`
	ResponseSuccess *string `json:"response_success"`
	ResponseError   *string `json:"response_error"`
	Used            *bool   `json:"used"`
}

type FlowData struct {
	Default string        `json:"default"`
	Name    string        `json:"name"`
	Actions []*FlowAction `json:"actions"`
}

type FlowAction struct {
	Used     bool    `json:"used"`
	Order    int     `json:"order"`
	Name     *string `json:"name,omitempty"`
	Current  string  `json:"current"`
	Default  string  `json:"default,omitempty"`
	Resposta string  `json:"resposta"`
	Error    string  `json:"error"`
	Type     string  `json:"type"`
}
