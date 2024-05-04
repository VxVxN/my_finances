package httptools

import (
	"encoding/json"
	"io"
)

func UnmarshalRequest(body io.ReadCloser, reqStruct interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(&reqStruct)
}
