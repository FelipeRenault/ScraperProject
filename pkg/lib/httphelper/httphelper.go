package httphelper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BadRequest(w http.ResponseWriter, body string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(body))
}

func InternalError(w http.ResponseWriter, body string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(body))
}

func RenderJSON(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-type", "application/json; charset=utf-8")

	if s, ok := v.(string); ok {
		fmt.Fprint(w, s)
		return nil
	}

	return json.NewEncoder(w).Encode(v)
}
