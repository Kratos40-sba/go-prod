package main

import "log"

type App struct {
}

func (app *App) Run() error {
	log.Println("App starting")
	return nil
}
func main() {
	log.Println("Hello from main function")
	app := App{}
	if err := app.Run(); err != nil {
		log.Println("Error while starting The APP : ", err)
	}
}
