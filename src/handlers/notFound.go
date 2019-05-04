package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NickTaporuk/redeam/src/utils"
)

type NotFoundJsonResponse struct {
	Message   string `json:"message"`
	Status    int    `json:"status"`
	Container string `json:"container"`
	Path      string `json:"path"`
	Time      int64  `json:"timestamp"`
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJson(w)

	notFoundResponse := &NotFoundJsonResponse{
		Message:   "Path not found",
		Status:    http.StatusNotFound,
		Container: r.Host,
		Path:      r.URL.Path,
		Time:      time.Now().UTC().Unix(),
	}

	json, _ := json.Marshal(notFoundResponse)
	w.Write(json)
}
