package dto

type OutputGetModelFlow struct {
	Default  interface{}   `json:"default"`
	FlowName string        `json:"name"`
	Flows    []interface{} `json:"flows"`
}
