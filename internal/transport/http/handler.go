package http

import (
	"encoding/json"
	"github.com/Kratos40-sba/go-prod/internal/comment"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Handler - stores pointer to o/ur comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object tio store responses from our API
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, request *http.Request) {
		if err := sendOkResponse(writer, Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods(http.MethodDelete)
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods(http.MethodPut)
}

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "cannot parse comment id", err)
		return
	}
	cmt, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(writer, "cannot get comment", err)
		return
	}
	if err := sendOkResponse(writer, cmt); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes comment by ID
func (h *Handler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	vars := mux.Vars(request)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "cannot parse id", err)
		return
	}
	if err := h.Service.DeleteComment(uint(i)); err != nil {
		sendErrorResponse(writer, "cannot delete comment", err)
		return
	}
	if err := sendOkResponse(writer, Response{Message: "Comment was deleted"}); err != nil {
		panic(err)
	}
}

// UpdateComment - update comment by id and a newComment
func (h *Handler) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	var updateComment comment.Comment
	if err := json.NewDecoder(request.Body).Decode(&updateComment); err != nil {
		sendErrorResponse(writer, "cannot parse request", err)
		return
	}
	vars := mux.Vars(request)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "cannot parse comment id", err)
		return
	}
	cmt, err := h.Service.UpdateComment(uint(i), updateComment)
	if err != nil {
		sendErrorResponse(writer, "cannot update comment", err)
		return
	}
	if err := sendOkResponse(writer, cmt); err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment
func (h *Handler) PostComment(writer http.ResponseWriter, request *http.Request) {
	var createdComment comment.Comment
	if err := json.NewDecoder(request.Body).Decode(&createdComment); err != nil {
		sendErrorResponse(writer, "cannot parse request", err)
		return
	}
	cmt, err := h.Service.PostComment(createdComment)
	if err != nil {
		sendErrorResponse(writer, "cannot create comment", err)
		return
	}
	if err := sendOkResponse(writer, cmt); err != nil {
		panic(err)
	}
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(writer http.ResponseWriter, request *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(writer, "cannot get all comments", err)
		return
	}
	if err := sendOkResponse(writer, comments); err != nil {
		panic(err)
	}
}
func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}
func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{
		Message: message,
		Error:   err.Error(),
	}); err != nil {
		panic(err)
	}
}
