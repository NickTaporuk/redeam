package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// initBooksTests is helper for init tests
//nolint
func initNotFoundHandlerTests(t *testing.T, method string, url string) (*http.Request, http.ResponseWriter) {
	var err error

	var wr = httptest.NewRecorder()

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	return req, wr
}
func TestNotFoundHandler(t *testing.T) {
	req, wr := initNotFoundHandlerTests(t, http.MethodGet, "/test/test/1")

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1 positive",
			args: args{
				w: wr,
				r: req,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotFoundHandler(tt.args.w, tt.args.r)
		})
	}
}
