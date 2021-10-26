package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Handler - stores pointer to o/ur comment service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "I am alive")
	})

}
