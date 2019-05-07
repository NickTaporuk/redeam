package app

import (
	"context"
	"log"
	"net/http"

	"github.com/NickTaporuk/redeam/src/configuration"
	"github.com/NickTaporuk/redeam/src/core"
	"github.com/NickTaporuk/redeam/src/db"
	"github.com/NickTaporuk/redeam/src/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type (
	// Runner
	Runner interface {
		Run() error
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
	// RunnerDatabase
	RunnerConfiguration interface {
		Configuration() *configuration.Config
		SetConfiguration(c *configuration.Config)
	}
	// Main structure use for run app
	Main struct {
		version string
		router  *mux.Router
		db      *gorm.DB
		config  *configuration.Config
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
func (m *Main) SetDB(d *gorm.DB) {
	m.db = d
}

// Configuration is app configuration getter
func (m *Main) Configuration() *configuration.Config {
	return m.config
}

// Configuration is app configuration getter
func (m *Main) SetConfiguration(c *configuration.Config) {
	m.config = c
}

// Main initialize with predefined configuration
func (m *Main) Init() error {
	var err error
	var conn *gorm.DB
	var conf *configuration.Config
	var dbCnf *configuration.DatabaseConfig
	var servCnf *configuration.ServiceConfig
	var data = make(map[string]string)
	var router *mux.Router
	// extract configuration data from env
	err = configuration.InitEnv(data)

	if err != nil {
		return err
	}
	// initiate db configuration
	dbCnf, err = configuration.NewDatabaseConfig(data)

	if err != nil {
		return err
	}

	servCnf = configuration.NewServiceConfig(data)

	conf = configuration.NewConfig(dbCnf, servCnf)

	conn, err = db.Init(conf)

	if err != nil {
		return err
	}
	// initiate routes
	router = mux.NewRouter()
	router.Use(m.AddContext)
	m.SetRouter(router)
	// initiate database connection
	m.SetDB(conn)
	// initiate not found handler
	m.Router().NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	// initiate database connection
	m.SetConfiguration(conf)
	// initiate routers scope
	m.SetRouters()

	return nil
}

// Run is base runner of application
func (m *Main) Run() error {

	log.Fatal(http.ListenAndServe(m.Configuration().ServicePort, m.Router()))

	return nil
}

// SetRouters initiate all endpoints
func (m *Main) SetRouters() {

	var r *mux.Router

	r = m.Router()

	r.HandleFunc(handlers.BaseRouteBooksName, handlers.GetBooks).Methods(http.MethodGet)
	r.HandleFunc(handlers.BaseRouteBooksNameByID, handlers.GetBook).Methods(http.MethodGet)
	r.HandleFunc(handlers.BaseRouteBooksName, handlers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc(handlers.BaseRouteBooksNameByID, handlers.UpdateBook).Methods(http.MethodPatch)
	r.HandleFunc(handlers.BaseRouteBooksNameByID, handlers.DeleteBook).Methods(http.MethodDelete)
}

// Close turn off data
func (m *Main) Close() (err error) {

	err = m.DB().Close()
	if err != nil {
		return err
	}

	return nil
}

func (m *Main) AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if m != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), core.ContextDbName, m.DB())
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
