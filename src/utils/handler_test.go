package utils

import (
	"net/http"
	"testing"
)

func TestRespondJSON(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		status  int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RespondJSON(tt.args.w, tt.args.status, tt.args.payload)
		})
	}
}

func TestRespondError(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		r       *http.Request
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RespondError(tt.args.w, tt.args.r, tt.args.code, tt.args.message)
		})
	}
}
