package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NickTaporuk/redeam/src/models"
	"github.com/NickTaporuk/redeam/src/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const (
	BaseRouteBooksName     = "/books"
	BaseRouteBooksNameById = "/books/{id}"
)

func GetBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJson(w)
	params := mux.Vars(r)
	fmt.Println(params["id"])
	var book models.Books
	book.ID = 1

	fmt.Println(db.Select(&book))
	fmt.Printf("%#v", &book)
	//for _, item := range books {
	//	if item.ID == params["id"] {
	//		json.NewEncoder(w).Encode(item)
	//		return
	//	}
	//}
	json.NewEncoder(w).Encode(&models.Books{})
}

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJson(w)
	var book models.Books
	//_ = json.NewDecoder(r.Body).Decode(&book)
	//book.ID = strconv.Itoa(rand.Intn(1000000))
	//books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJson(w)
	//params := mux.Vars(r)
	//for index, item := range books {
	//	if item.ID == params["id"] {
	//		books = append(books[:index], books[index+1:]...)
	//		var book Book
	//		_ = json.NewDecoder(r.Body).Decode(&book)
	//		book.ID = params["id"]
	//		books = append(books, book)
	//		json.NewEncoder(w).Encode(book)
	//		return
	//	}
	//}
	json.NewEncoder(w).Encode(nil)
}

func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	utils.AddHeaderContentTypeJson(w)
	//params := mux.Vars(r)
	//for index, item := range books {
	//	if item.ID == params["id"] {
	//		books = append(books[:index], books[index+1:]...)
	//		break
	//	}
	//}
	json.NewEncoder(w).Encode(nil)
}
