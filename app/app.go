package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/opAPIProgression/app/config"
	"github.com/opAPIProgression/app/handler"
	"github.com/opAPIProgression/app/model"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/bosses", a.GetAllBosses)
	a.Post("/bosses", a.CreateBoss)
	a.Get("/bosses/{title}", a.GetBoss)
	a.Put("/bosses/{title}", a.UpdateBoss)
	a.Delete("/bosses/{title}", a.DeleteBoss)
	a.Put("/bosses/{title}/killboss", a.KillBoss)
	a.Put("/bosses/{title}/reviveboss", a.ReviveBoss)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
func (a *App) GetAllBosses(w http.ResponseWriter, r *http.Request) {
	handler.GetAllBosses(a.DB, w, r)
}

func (a *App) CreateBoss(w http.ResponseWriter, r *http.Request) {
	handler.CreateBoss(a.DB, w, r)
}

func (a *App) GetBoss(w http.ResponseWriter, r *http.Request) {
	handler.GetBoss(a.DB, w, r)
}

func (a *App) UpdateBoss(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBoss(a.DB, w, r)
}

func (a *App) DeleteBoss(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBoss(a.DB, w, r)
}

func (a *App) KillBoss(w http.ResponseWriter, r *http.Request) {
	handler.KillBoss(a.DB, w, r)
}

func (a *App) ReviveBoss(w http.ResponseWriter, r *http.Request) {
	handler.ReviveBoss(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
