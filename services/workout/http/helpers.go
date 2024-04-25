package http

import (
	"encoding/json"
	"io"
)

func readJSON(r io.ReadCloser, data any) error {
	defer r.Close()
	return json.NewDecoder(r).Decode(data)
}

func writeJSON(w io.Writer, data any) error {
	return json.NewEncoder(w).Encode(data)
}
