package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	router *mux.Router
	db     *gorm.DB
}

func NewApp() App {
	var err error
	app := App{}
	app.router = mux.NewRouter()
	app.db, err = gorm.Open(sqlite.Open("phoenix.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	app.db.AutoMigrate(&User{}, &Music{})
	return app
}

func main() {
	app := NewApp()
	app.SetupRoutes()
	router := handlers.LoggingHandler(os.Stdout, app.router)
	http.Handle("/", handlers.CompressHandler(router))
	fmt.Println("starting the server...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
