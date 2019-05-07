package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	ErrorJSONResponse struct {
		Message   string `json:"message"`
		Status    int    `json:"status"`
		Container string `json:"container"`
		Path      string `json:"path"`
		Time      int64  `json:"timestamp"`
	}
)

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	AddHeaderContentTypeJSON(w)

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, r *http.Request, code int, message string) {

	errorResponse := &ErrorJSONResponse{
		Message:   message,
		Status:    code,
		Container: r.Host,
		Path:      r.URL.Path,
		Time:      time.Now().UTC().Unix(),
	}

	RespondJSON(w, code, errorResponse)
}
