package main

import (
	"github.com/Kratos40-sba/go-prod/internal/comment"
	"github.com/Kratos40-sba/go-prod/internal/database"
	transportHTTP "github.com/Kratos40-sba/go-prod/internal/transport/http"
	"log"
	"net/http"
)

// App - contains a pointer to db connection
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db)
	handler := transportHTTP.NewHandler(commentService)
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
