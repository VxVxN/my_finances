package httptools

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	data, err := json.Marshal(errorResponse{Error: err.Error()})
	if err != nil {
		// todo
		return
	}
	w.Write(data)
}

func SuccessResponse(w http.ResponseWriter, resp interface{}) {
	w.WriteHeader(http.StatusOK)
	if resp == nil {
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		// todo
		return
	}
	w.Write(data)
}
