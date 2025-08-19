package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, msg map[string]interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msg)
}

func WriteJSONError(w http.ResponseWriter, status int, err error) {
	errMap := make(map[string]interface{})
	errMap["error"] = err
	WriteJSONResponse(w, status, errMap)
}
