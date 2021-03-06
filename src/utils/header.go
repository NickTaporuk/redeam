package utils

import "net/http"

const (
	HeaderContentTypeKey  = "Content-Type"
	HeaderContentTypeJSON = "application/json"
)

// AddHeader is wrapper for add header key value pair
func AddHeader(w http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}

// AddHeader is wrapper for add header key value pair
func AddHeaderContentTypeJSON(w http.ResponseWriter) {
	AddHeader(w, HeaderContentTypeKey, HeaderContentTypeJSON)
}
