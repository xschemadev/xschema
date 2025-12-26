package retriever

import "encoding/json"

type Schema struct {
	Name   string          `json:"name"`
	Schema json.RawMessage `json:"schema"`
}

// RetrieveFromURL fetches a JSON schema from a URL
func RetrieveFromURL(url string) (json.RawMessage, error) {
	// TODO: implement
	return nil, nil
}

// RetrieveFromFile reads a JSON schema from a file
func RetrieveFromFile(path string) (json.RawMessage, error) {
	// TODO: implement
	return nil, nil
}
