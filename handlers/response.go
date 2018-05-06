package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response represents a JSON response that a handler would use to output
// JSON to the client per request
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// Marshall a response and outputs it to the writer
func JsonWriter(w http.ResponseWriter, res *Response) {
	b, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{error: "%s"}`, err)))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(res.Code)
	w.Write(b)
}
