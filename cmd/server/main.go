package main

import (
	transportHTTP "github.com/Kratos40-sba/go-prod/internal/transport/http"
	"log"
	"net/http"
)

// App - contains a pointer to db connection
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Println("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	log.Println("Hello from main function")
	app := App{}
	if err := app.Run(); err != nil {
		log.Println("Error while starting The APP : ", err)
	}

}
