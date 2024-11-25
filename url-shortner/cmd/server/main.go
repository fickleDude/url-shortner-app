package main

import (
	"log"
	"net/http"

	db "example.com/url-shortner/db"
	router "example.com/url-shortner/router"
	service "example.com/url-shortner/service"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models service.Models
}

func (app *Application) Serve() error {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Routes(),
	}
	return srv.ListenAndServe()
}

var connString = "postgres://postgres:postgres@localhost:5432/demo?sslmode=disable"

func main() {
	dbConn, err := db.ConnectPostgres(connString)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	cfg := Config{
		Port: "8080",
	}

	defer dbConn.DB.Close()
	app := &Application{
		Config: cfg,
		Models: service.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
