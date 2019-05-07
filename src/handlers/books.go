package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/NickTaporuk/redeam/src/core"
	"github.com/NickTaporuk/redeam/src/models"
	"github.com/NickTaporuk/redeam/src/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const (
	BaseRouteBooksName     = "/api/v1/books"
	BaseRouteBooksNameByID = "/api/v1/books/{id}"
)

var (
	// ErrorRecordByIDNotFound is error when record isn't found in database table
	ErrorRecordByIDNotFound = errors.New("record was not found by ID")
	// ErrorHTTPParameterIDRequired is error when record isn't found in database table
	ErrorHTTPParameterIDRequired = errors.New("parameter ID is required")
	// ErrorDatabaseNotInitiate is error when db isn't in request context
	ErrorDatabaseNotInitiate = errors.New("database isn't initialized by context")
	// ErrorHTTPBodyEmpty is error when http body is empty
	ErrorHTTPBodyEmpty = errors.New("http body is empty")
)

// GetBooks is handler for book rest api
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Books
	var err error
	var db *gorm.DB
	var ctxDb interface{}

	books = []models.Books{}

	ctxDb = r.Context().Value(core.ContextDbName)
	if ctxDb == nil {
		err = ErrorDatabaseNotInitiate
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	db = ctxDb.(*gorm.DB)

	if err = db.Find(&books).Error; gorm.IsRecordNotFoundError(err) {
		utils.RespondError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, books)

}

// GetBook is handler for book rest api
func GetBook(w http.ResponseWriter, r *http.Request) {
	var bookID int
	var bookIDConverted uint64
	var err error
	var db *gorm.DB

	var ctxDb = r.Context().Value(core.ContextDbName)
	if ctxDb == nil {
		err = ErrorDatabaseNotInitiate
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	db = ctxDb.(*gorm.DB)
	utils.AddHeaderContentTypeJSON(w)
	params := mux.Vars(r)

	if params["id"] != "" {
		bookID, err = strconv.Atoi(params["id"])
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		var book models.Books
		bookIDConverted = uint64(bookID)

		book.ID = bookIDConverted

		if db.First(&book).RecordNotFound() {
			err = ErrorRecordByIDNotFound
			utils.RespondError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		utils.RespondJSON(w, http.StatusOK, book)

	} else {
		err = ErrorHTTPParameterIDRequired
		utils.RespondError(w, r, http.StatusBadRequest, err.Error())
		return
	}
}

// CreateBook is handler for book rest api
func CreateBook(w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJSON(w)
	var book models.Books
	var err error
	var db *gorm.DB

	var ctxDb = r.Context().Value(core.ContextDbName)

	if ctxDb == nil {
		err = ErrorDatabaseNotInitiate
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	db = ctxDb.(*gorm.DB)

	if r.Body == nil {
		err = ErrorHTTPBodyEmpty
		utils.RespondError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err = db.Create(&book).Error; err != nil {
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, book)
}

// UpdateBook is handler for book rest api
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var err error
	var db *gorm.DB
	var bookID int
	var bookIDConverted uint64

	var ctxDb = r.Context().Value(core.ContextDbName)
	if ctxDb == nil {
		err = ErrorDatabaseNotInitiate
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	db = ctxDb.(*gorm.DB)

	utils.AddHeaderContentTypeJSON(w)
	params := mux.Vars(r)

	if params["id"] != "" {

		var book models.Books

		if r.Body == nil {
			err = ErrorHTTPBodyEmpty
			utils.RespondError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		bookID, err = strconv.Atoi(params["id"])
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		bookIDConverted = uint64(bookID)
		book.ID = bookIDConverted
		book.UpdatedAt = time.Now()

		var data = make(map[string]interface{})

		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		if err = db.Model(&book).Updates(data).Error; err != nil {
			utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondJSON(w, http.StatusOK, &book)

	} else {
		err = ErrorHTTPParameterIDRequired
		utils.RespondError(w, r, http.StatusBadRequest, err.Error())
		return
	}
}

// DeleteBook is handler for book rest api
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var bookID int
	var bookIDConverted uint64
	var err error
	var db *gorm.DB

	var ctxDb = r.Context().Value(core.ContextDbName)
	if ctxDb == nil {
		err = ErrorDatabaseNotInitiate
		utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	db = ctxDb.(*gorm.DB)

	utils.AddHeaderContentTypeJSON(w)
	params := mux.Vars(r)

	if params["id"] != "" {
		bookID, err = strconv.Atoi(params["id"])
		if err != nil {
			utils.RespondError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		var book models.Books
		bookIDConverted = uint64(bookID)

		book.ID = bookIDConverted

		if db.First(&book).RecordNotFound() {
			err = ErrorRecordByIDNotFound
			utils.RespondError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		db.Model(&book).Delete(book)

		utils.RespondJSON(w, http.StatusOK, struct {
			Message string
		}{
			Message: "record deleted",
		})

	} else {
		err = ErrorHTTPParameterIDRequired
		utils.RespondError(w, r, http.StatusBadRequest, err.Error())
		return
	}

}
