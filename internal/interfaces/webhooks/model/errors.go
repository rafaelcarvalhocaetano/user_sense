package model

// Errors contém informações sobre erros ocorridos.
type Error struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details,omitempty"`
	Href    string `json:"href,omitempty"`
}
