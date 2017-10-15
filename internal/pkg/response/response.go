package response

import (
	"botwhatsapp/internal/infra/ports"
	"encoding/json"
	"github.com/k0kubun/pp/v3"
	"net/http"
)

type Payload struct {
	Output any `json:"output"`
	Input  any `json:"input"`
}

type Response struct {
	log    ports.Logger
	module string
}

func NewResponse(m string, log ports.Logger) *Response {
	return &Response{log: log, module: m}
}

func (e *Response) Error(input, msg any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	payload := Payload{Output: msg, Input: input}
	e.log.Error(e.module, payload, http.StatusBadRequest)
	output, _ := json.Marshal(payload)
	_, _ = pp.Println(msg)

	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(output)
}

func (e *Response) Success(input, msg any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var x struct {
		Data any `json:"data"`
	}
	x.Data = msg
	e.log.Debug(e.module, Payload{Input: input, Output: msg}, http.StatusOK)
	output, _ := json.Marshal(x)
	_, _ = pp.Println(msg)

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(output)
}
