package handlers

import (
	"encoding/json"		// For JSON encoding
	"net/http"			//HTTP server and request handling
)

func jsonRes(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
