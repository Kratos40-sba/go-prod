package main

import (
	"github.com/Kratos40-sba/go-prod/internal/comment"
	"github.com/Kratos40-sba/go-prod/internal/database"
	transportHTTP "github.com/Kratos40-sba/go-prod/internal/transport/http"
	logLib "github.com/sirupsen/logrus"
	"net/http"
)

// App - contains a pointer to db connection
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	logLib.SetFormatter(&logLib.JSONFormatter{})
	logLib.WithFields(logLib.Fields{
		"AppName":    app.Name,
		"AppVersion": app.Version,
	}).Info("Setting up application")
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
		logLib.Error("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name:    "Commenting Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		logLib.Error("Error starting up our REST API")
		logLib.Fatal(err)
	}

}
