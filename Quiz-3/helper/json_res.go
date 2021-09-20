package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface {}, status int) {
	w.Header().Set("content-type", "application/json")
	res, _ := json.Marshal(p)

	w.WriteHeader(status)
	w.Write([]byte(res))
}