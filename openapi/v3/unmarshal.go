package v3

import (
	"encoding/json"
	"io"
)

func Unmarshal(r io.Reader) (*OpenAPI, error) {
	var spec OpenAPI
	if err := json.NewDecoder(r).Decode(&spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
