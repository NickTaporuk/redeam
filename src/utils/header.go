package utils

import "net/http"

const (
	HeaderContentTypeKey  = "Content-Type"
	HeaderContentTypeJson = "application/json"
)

// AddHeader is wrapper for add header key value pair
func AddHeader(w http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}

// AddHeader is wrapper for add header key value pair
func AddHeaderContentTypeJson(w http.ResponseWriter) {
	AddHeader(w, HeaderContentTypeKey, HeaderContentTypeJson)
}
