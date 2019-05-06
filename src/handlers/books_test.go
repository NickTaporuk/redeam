package handlers

import (
	"bytes"
	"context"
	_ "database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NickTaporuk/redeam/src/core"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

const (
	dbType                     = "postgres"
	TestBaseRouteBooksNameByID = "/api/v1/books/%d"
)

// initBooksTests is helper for init tests
func initBooksTests(t *testing.T, method string, URL string, httpVars map[string]string, body io.Reader, ctxAdd bool) (sqlmock.Sqlmock, *http.Request, http.ResponseWriter, error) {
	var err error
	var gormDB *gorm.DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err = gorm.Open(dbType, db)
	gormDB.LogMode(true)

	var wr *httptest.ResponseRecorder
	wr = httptest.NewRecorder()

	req, err := http.NewRequest(method, URL, body)

	if httpVars != nil {
		req = mux.SetURLVars(req, httpVars)
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	if ctxAdd {
		ctx := context.WithValue(req.Context(), core.ContextDbName, gormDB)
		req = req.WithContext(ctx)
	}

	return mock, req, wr, nil
}

func TestGetBooks(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "author"}).AddRow(1, "Jack")
	rows1 := sqlmock.NewRows([]string{"id", "author"})

	mock, req, wr, err := initBooksTests(t, http.MethodGet, BaseRouteBooksName, nil, nil, true)
	assert.NoError(t, err)
	mock1, req1, wr1, err1 := initBooksTests(t, http.MethodGet, BaseRouteBooksName, nil, nil, true)
	assert.NoError(t, err1)
	mock2, req2, wr2, err2 := initBooksTests(t, http.MethodGet, BaseRouteBooksName, nil, nil, false)
	assert.NoError(t, err2)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	var tests = []struct {
		name         string
		args         args
		mockDB       sqlmock.Sqlmock
		responseCode int
		rows         *sqlmock.Rows
		err          error
		wantErr      bool
	}{
		{
			name: "GetBooks handler test 1 positive",
			args: args{
				r: req,
				w: wr,
			},
			mockDB:       mock,
			responseCode: http.StatusOK,
			rows:         rows,
		},
		{
			name: "GetBooks handler test 2 negative emulate db error",
			args: args{
				r: req1,
				w: wr1,
			},
			mockDB:       mock1,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows1,
			err:          gorm.ErrRecordNotFound,
		},
		{
			name: "GetBooks handler test 3 negative without context",
			args: args{
				r: req2,
				w: wr2,
			},
			mockDB:       mock2,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			err:          nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var w *httptest.ResponseRecorder

			tt.mockDB.
				ExpectQuery("^SELECT (.+) FROM \"books\"").
				WillReturnRows(tt.rows).
				WillReturnError(tt.err)

			GetBooks(tt.args.w, tt.args.r)
			w = tt.args.w.(*httptest.ResponseRecorder)

			if status := w.Code; status != tt.responseCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseCode)
			}

			// we make sure that all expectations were met
			if err := tt.mockDB.ExpectationsWereMet(); tt.err != nil && err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	var rows, rows1 *sqlmock.Rows

	rows = sqlmock.NewRows([]string{"id", "author"}).AddRow(1, "Jack")
	rows1 = sqlmock.NewRows([]string{"id", "author"})

	var path string
	var vars = map[string]string{
		"id": "1",
	}
	path = fmt.Sprintf(TestBaseRouteBooksNameByID, 1)

	mock, req, wr, err := initBooksTests(t, http.MethodGet, path, vars, nil, true)
	assert.NoError(t, err)
	mock1, req1, wr1, err1 := initBooksTests(t, http.MethodGet, path, vars, nil, true)
	assert.NoError(t, err1)
	mock2, req2, wr2, err2 := initBooksTests(t, http.MethodGet, path, nil, nil, false)
	assert.NoError(t, err2)
	mock3, req3, wr3, err3 := initBooksTests(t, http.MethodGet, path, nil, nil, true)
	assert.NoError(t, err3)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name         string
		args         args
		mockDB       sqlmock.Sqlmock
		responseCode int
		rows         *sqlmock.Rows
		err          error
		wantErr      bool
	}{
		{
			name: "GetBook handler test 1 positive",
			args: args{
				r: req,
				w: wr,
			},
			mockDB:       mock,
			responseCode: http.StatusOK,
			rows:         rows,
		},
		{
			name: "GetBook handler test 2 negative emulate db error",
			args: args{
				r: req1,
				w: wr1,
			},
			mockDB:       mock1,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows1,
			err:          gorm.ErrRecordNotFound,
		},
		{
			name: "GetBook handler test 3 negative without context",
			args: args{
				r: req2,
				w: wr2,
			},
			mockDB:       mock2,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			err:          nil,
		},
		{
			name: "GetBook handler test 4 negative without params",
			args: args{
				r: req3,
				w: wr3,
			},
			mockDB:       mock3,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			err:          nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w *httptest.ResponseRecorder

			tt.mockDB.
				ExpectQuery("^SELECT (.+) FROM \"books\" WHERE \"books\".\"id\" = (.+) ORDER BY \"books\".\"id\" ASC LIMIT 1").
				WillReturnRows(tt.rows).
				WillReturnError(tt.err)
			handler := http.HandlerFunc(GetBook)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			w = tt.args.w.(*httptest.ResponseRecorder)

			if status := w.Code; status != tt.responseCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseCode)
			}

			// we make sure that all expectations were met
			if err := tt.mockDB.ExpectationsWereMet(); tt.err != nil && err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	var rows, rows1, rows2 *sqlmock.Rows

	rows = sqlmock.
		NewRows([]string{"title", "author", "publisher", "publish_date", "rating", "status", "created_at", "updated_at"}).
		AddRow("test title", "test author", "test publisher", "0001-01-01 00:00:00", 0, false, "2019-05-06 14:53:47", "2019-05-06 14:53:47")

	rows1 = sqlmock.NewRows([]string{"id", "author"})
	// rows2 is sql result set
	rows2 = sqlmock.NewRows([]string{"id"}).AddRow("1")

	var path string
	var vars = map[string]string{
		"id": "1",
	}

	payload := []byte(`{"author":"test author","publisher":"test publisher","title":"test title"}`)
	payload1 := []byte(`    `)

	path = fmt.Sprintf(TestBaseRouteBooksNameByID, 1)

	mock, req, wr, err := initBooksTests(t, http.MethodPost, path, vars, bytes.NewBuffer(payload), true)
	assert.NoError(t, err)

	mock1, req1, wr1, err1 := initBooksTests(t, http.MethodPost, path, vars, bytes.NewBuffer(payload), true)
	assert.NoError(t, err1)

	mock2, req2, wr2, err2 := initBooksTests(t, http.MethodPost, path, nil, bytes.NewBuffer(payload), false)
	assert.NoError(t, err2)

	mock3, req3, wr3, err3 := initBooksTests(t, http.MethodPost, path, nil, bytes.NewBuffer(payload), true)
	assert.NoError(t, err3)

	mock4, req4, wr4, err4 := initBooksTests(t, http.MethodPost, path, nil, nil, true)
	assert.NoError(t, err4)

	mock5, req5, wr5, err5 := initBooksTests(t, http.MethodPost, path, nil, bytes.NewBuffer(payload1), true)
	assert.NoError(t, err5)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name         string
		args         args
		mockDB       sqlmock.Sqlmock
		responseCode int
		rows         *sqlmock.Rows
		resultRows   *sqlmock.Rows
		err          error
		wantErr      bool
	}{
		{
			name: "CreateBook handler test 1 positive",
			args: args{
				r: req,
				w: wr,
			},
			mockDB:       mock,
			responseCode: http.StatusCreated,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "CreateBook handler test 2 negative emulate db error",
			args: args{
				r: req1,
				w: wr1,
			},
			mockDB:       mock1,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows1,
			resultRows:   rows2,
			err:          gorm.ErrRecordNotFound,
		},
		{
			name: "CreateBook handler test 3 negative without context",
			args: args{
				r: req2,
				w: wr2,
			},
			mockDB:       mock2,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "CreateBook handler test 4 negative without params",
			args: args{
				r: req3,
				w: wr3,
			},
			mockDB:       mock3,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "CreateBook handler test 5 negative without body",
			args: args{
				r: req4,
				w: wr4,
			},
			mockDB:       mock4,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "CreateBook handler test 6 negative with empty body",
			args: args{
				r: req5,
				w: wr5,
			},
			mockDB:       mock5,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w *httptest.ResponseRecorder

			tt.mockDB.
				ExpectQuery("^INSERT INTO \"books\" (.+)").
				WillReturnRows(tt.resultRows).
				WillReturnError(tt.err)

			handler := http.HandlerFunc(CreateBook)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			w = tt.args.w.(*httptest.ResponseRecorder)

			if status := w.Code; status != tt.responseCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseCode)
			}

			// we make sure that all expectations were met
			if err := tt.mockDB.ExpectationsWereMet(); tt.err != nil && err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	var rows, rows1, rows2 *sqlmock.Rows

	rows = sqlmock.
		NewRows([]string{"author"}).
		AddRow("test author")

	rows1 = sqlmock.NewRows([]string{"author"})
	// rows2 is sql result set
	rows2 = sqlmock.NewRows([]string{"author"}).AddRow("test")

	var path string
	var vars = map[string]string{
		"id": "1",
	}

	payload := []byte(`{"author":"test"}`)
	payload1 := []byte(`    `)

	path = fmt.Sprintf(TestBaseRouteBooksNameByID, 1)

	mock, req, wr, err := initBooksTests(t, http.MethodPatch, path, vars, bytes.NewBuffer(payload), true)
	assert.NoError(t, err)

	mock1, req1, wr1, err1 := initBooksTests(t, http.MethodPatch, path, vars, bytes.NewBuffer(payload), true)
	assert.NoError(t, err1)

	mock2, req2, wr2, err2 := initBooksTests(t, http.MethodPatch, path, nil, bytes.NewBuffer(payload), false)
	assert.NoError(t, err2)

	mock3, req3, wr3, err3 := initBooksTests(t, http.MethodPatch, path, nil, bytes.NewBuffer(payload), true)
	assert.NoError(t, err3)

	mock4, req4, wr4, err4 := initBooksTests(t, http.MethodPatch, path, nil, nil, true)
	assert.NoError(t, err4)

	mock5, req5, wr5, err5 := initBooksTests(t, http.MethodPatch, path, nil, bytes.NewBuffer(payload1), true)
	assert.NoError(t, err5)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name         string
		args         args
		mockDB       sqlmock.Sqlmock
		responseCode int
		rows         *sqlmock.Rows
		resultRows   *sqlmock.Rows
		err          error
		wantErr      bool
	}{
		{
			name: "UpdateBook handler test 1 positive",
			args: args{
				r: req,
				w: wr,
			},
			mockDB:       mock,
			responseCode: http.StatusOK,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "UpdateBook handler test 2 negative emulate db error",
			args: args{
				r: req1,
				w: wr1,
			},
			mockDB:       mock1,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows1,
			resultRows:   rows2,
			err:          gorm.ErrRecordNotFound,
		},
		{
			name: "UpdateBook handler test 3 negative without context",
			args: args{
				r: req2,
				w: wr2,
			},
			mockDB:       mock2,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "UpdateBook handler test 4 negative without params",
			args: args{
				r: req3,
				w: wr3,
			},
			mockDB:       mock3,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "UpdateBook handler test 5 negative without body",
			args: args{
				r: req4,
				w: wr4,
			},
			mockDB:       mock4,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "UpdateBook handler test 6 negative with empty body",
			args: args{
				r: req5,
				w: wr5,
			},
			mockDB:       mock5,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w *httptest.ResponseRecorder

			tt.mockDB.
				ExpectExec("^UPDATE \"books\" SET \"author\" = (.+), \"updated_at\" = (.+)  WHERE \"books\".\"id\" = (.+)").
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.err)

			handler := http.HandlerFunc(UpdateBook)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			w = tt.args.w.(*httptest.ResponseRecorder)

			if status := w.Code; status != tt.responseCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseCode)
			}

			// we make sure that all expectations were met
			if err := tt.mockDB.ExpectationsWereMet(); tt.err != nil && err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	var rows, rows1, rows2 *sqlmock.Rows

	rows = sqlmock.
		NewRows([]string{"author"}).
		AddRow("test author")

	rows1 = sqlmock.NewRows([]string{"author"})
	// rows2 is sql result set
	rows2 = sqlmock.NewRows([]string{"author"}).AddRow("test")

	var path string
	var vars = map[string]string{
		"id": "1",
	}

	payload := []byte(`{"author":"test"}`)
	payload1 := []byte(`    `)

	path = fmt.Sprintf(TestBaseRouteBooksNameByID, 1)

	mock, req, wr, err := initBooksTests(t, http.MethodPatch, path, vars, bytes.NewBuffer(payload), true)
	assert.NoError(t, err)

	mock1, req1, wr1, err1 := initBooksTests(t, http.MethodPatch, path, vars, bytes.NewBuffer(payload), true)
	assert.NoError(t, err1)

	mock2, req2, wr2, err2 := initBooksTests(t, http.MethodPatch, path, nil, bytes.NewBuffer(payload), false)
	assert.NoError(t, err2)

	mock3, req3, wr3, err3 := initBooksTests(t, http.MethodPatch, path, nil, bytes.NewBuffer(payload), true)
	assert.NoError(t, err3)

	mock4, req4, wr4, err4 := initBooksTests(t, http.MethodPatch, path, nil, nil, true)
	assert.NoError(t, err4)

	mock5, req5, wr5, err5 := initBooksTests(t, http.MethodPatch, path, nil, bytes.NewBuffer(payload1), true)
	assert.NoError(t, err5)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name         string
		args         args
		mockDB       sqlmock.Sqlmock
		responseCode int
		rows         *sqlmock.Rows
		resultRows   *sqlmock.Rows
		err          error
		wantErr      bool
	}{
		{
			name: "DeleteBook handler test 1 positive",
			args: args{
				r: req,
				w: wr,
			},
			mockDB:       mock,
			responseCode: http.StatusOK,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "DeleteBook handler test 2 negative emulate db error",
			args: args{
				r: req1,
				w: wr1,
			},
			mockDB:       mock1,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows1,
			resultRows:   rows2,
			err:          gorm.ErrRecordNotFound,
		},
		{
			name: "DeleteBook handler test 3 negative without context",
			args: args{
				r: req2,
				w: wr2,
			},
			mockDB:       mock2,
			responseCode: http.StatusInternalServerError,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "DeleteBook handler test 4 negative without params",
			args: args{
				r: req3,
				w: wr3,
			},
			mockDB:       mock3,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "DeleteBook handler test 5 negative without body",
			args: args{
				r: req4,
				w: wr4,
			},
			mockDB:       mock4,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
		{
			name: "DeleteBook handler test 6 negative with empty body",
			args: args{
				r: req5,
				w: wr5,
			},
			mockDB:       mock5,
			responseCode: http.StatusBadRequest,
			wantErr:      true,
			rows:         rows,
			resultRows:   rows2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w *httptest.ResponseRecorder

			tt.mockDB.
				ExpectExec("^DELETE FROM \"books\"").
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.err)

			handler := http.HandlerFunc(DeleteBook)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			w = tt.args.w.(*httptest.ResponseRecorder)

			if status := w.Code; status != tt.responseCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseCode)
			}

			// we make sure that all expectations were met
			if err := tt.mockDB.ExpectationsWereMet(); tt.err != nil && err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
