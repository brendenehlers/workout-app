package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func readJSON(r io.ReadCloser, data any) error {
	defer r.Close()
	return json.NewDecoder(r).Decode(data)
}

func writeJSON(w http.ResponseWriter, data any) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
