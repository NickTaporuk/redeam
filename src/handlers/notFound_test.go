package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// initBooksTests is helper for init tests
func initNotFoundHandlerTests(t *testing.T, method string, URL string) (*http.Request, http.ResponseWriter, error) {
	var err error

	var wr *httptest.ResponseRecorder
	wr = httptest.NewRecorder()

	req, err := http.NewRequest(method, URL, nil)

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	return req, wr, nil
}
func TestNotFoundHandler(t *testing.T) {
	req, wr, err := initNotFoundHandlerTests(t, http.MethodGet, "/test/test/1")
	assert.NoError(t, err)

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
