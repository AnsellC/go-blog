package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ansellc/go-blog/database"
	"github.com/gorilla/mux"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Run listens for connections
func (a *App) Run(port string) {

}

// Init bootstraps the app
func (a *App) Init() {
	a.DB = database.Connect(os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_DBNAME"), os.Getenv("APP_DB_PORT"))
	a.Router = mux.NewRouter()
}

// Serve listen and serve the application using the port specified
func (a *App) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
