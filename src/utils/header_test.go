package utils

import (
	"net/http"
	"testing"
)

func TestAddHeader(t *testing.T) {
	type args struct {
		w     http.ResponseWriter
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddHeader(tt.args.w, tt.args.key, tt.args.value)
		})
	}
}

func TestAddHeaderContentTypeJSON(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddHeaderContentTypeJSON(tt.args.w)
		})
	}
}
