package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NickTaporuk/redeam/src/utils"
)

// NotFoundHandler is not found route handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJSON(w)

	notFoundResponse := &utils.ErrorJSONResponse{
		Message:   "Path not found",
		Status:    http.StatusNotFound,
		Container: r.Host,
		Path:      r.URL.Path,
		Time:      time.Now().UTC().Unix(),
	}

	data, _ := json.Marshal(notFoundResponse)
	w.Write(data)
}
