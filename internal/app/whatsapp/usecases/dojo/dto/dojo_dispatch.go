package dto

type DojoDispatchParams struct {
	Name   *string         `json:"name"`
	To     *string         `json:"to"`
	Next   *string         `json:"next"`
	Params *map[string]any `json:"params"`
}
