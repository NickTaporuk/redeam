package app

import (
	"log"
	"net/http"
	"os"

	"github.com/NickTaporuk/redeam/src/db"
	"github.com/NickTaporuk/redeam/src/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type (
	// Runner use for implement method Run
	Runner interface {
		Run() error
		Close()
	}
	// RunnerVersion
	RunnerVersion interface {
		Version() string
		SetVersion(v string)
	}
	// RunnerRouter
	RunnerRouter interface {
		Router() *mux.Router
		SetRouter(r *mux.Router)
	}
	// RunnerDatabase
	RunnerDatabase interface {
		DB() *gorm.DB
		SetDB(db *gorm.DB)
	}
	// Main structure use for run app
	Main struct {
		version string
		router  *mux.Router
		db      *gorm.DB
	}
)

// Version is version getter
func (m *Main) Version() string {
	return m.version
}

// SetVersion is Version setter
func (m *Main) SetVersion(v string) {
	m.version = v
}

// Router is router getter
func (m *Main) Router() *mux.Router {
	return m.router
}

// SetRouter is router setter
func (m *Main) SetRouter(r *mux.Router) {
	m.router = r
}

// DB is database getter
func (m *Main) DB() *gorm.DB {
	return m.db
}

// SetDB is database setter
func (m *Main) SetDB(db *gorm.DB) {
	m.db = db
}

// Main initialize with predefined configuration
func (m *Main) Init() error {
	var err error
	var conn *gorm.DB

	//TODO: will add check configuration
	// initiate db
	conn, err = db.Init()
	defer conn.Close()

	if err != nil {
		return err
	}
	// initiate routes
	m.SetRouter(mux.NewRouter())
	// initiate not found handler
	m.Router().NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	// initiate database connection
	m.SetDB(conn)
	// initiate routers scope
	m.SetRouters()

	return nil
}

// Run is base runner of application
func (m *Main) Run() error {
	var port string
	//TODO: move to configuration struct
	port = os.Getenv("APP_SERVER_PORT")

	log.Fatal(http.ListenAndServe(port, m.Router()))

	return nil
}

func (m *Main) SetRouters() {
	m.Get(handlers.BaseRouteBooksName, m.GetBooks)
	m.Get(handlers.BaseRouteBooksNameById, m.GetBooks)
	m.Post(handlers.BaseRouteBooksName, m.CreateBook)
	m.Put(handlers.BaseRouteBooksNameById, m.UpdateBook)
	m.Delete(handlers.BaseRouteBooksNameById, m.DeleteBook)
}

// Wrap the router for GET method
func (m *Main) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router().HandleFunc(path, f).Methods(http.MethodGet)
}

// Wrap the router for POST method
func (m *Main) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router().HandleFunc(path, f).Methods(http.MethodPost)
}

// Wrap the router for PUT method
func (m *Main) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router().HandleFunc(path, f).Methods(http.MethodPut)
}

// Wrap the router for PATCH method
func (m *Main) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router().HandleFunc(path, f).Methods(http.MethodPatch)
}

// Wrap the router for DELETE method
func (m *Main) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router().HandleFunc(path, f).Methods(http.MethodDelete)
}

//To manage Books Data method GET
func (m *Main) GetBooks(w http.ResponseWriter, r *http.Request) {
	handlers.GetBooks(m.DB(), w, r)
}

//To manage Books Data method POST
func (m *Main) CreateBook(w http.ResponseWriter, r *http.Request) {
	handlers.CreateBook(m.DB(), w, r)
}

//To manage Books Data method PUT
func (m *Main) UpdateBook(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateBook(m.DB(), w, r)
}

//To manage Books Data method DELETE
func (m *Main) DeleteBook(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteBook(m.DB(), w, r)
}

func (m *Main) Close() {}
