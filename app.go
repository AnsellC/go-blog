package main

import (
	"database/sql"

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
	a.DB = database.Connect("127.0.0.1", "root", "wisebread", "wisebread_compare", "3306")
	a.Router = mux.NewRouter()
}
