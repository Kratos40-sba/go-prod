package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Kratos40-sba/go-prod/internal/comment"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"os"
	"strings"

	logLib "github.com/sirupsen/logrus"
	"net/http"
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
	logLib.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)
	h.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, request *http.Request) {
		if err := sendOkResponse(writer, Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment", JwtAuth(h.PostComment)).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/{id}", JwtAuth(h.DeleteComment)).Methods(http.MethodDelete)
	h.Router.HandleFunc("/api/comment/{id}", JwtAuth(h.UpdateComment)).Methods(http.MethodPut)
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
		logLib.Error(err)
	}
}

// LoggingMiddleware - adds middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logLib.WithFields(
			logLib.Fields{
				"Method": request.Method,
				"Path":   request.URL.Path,
				"IP":     request.RemoteAddr,
			}).Info("Endpoint hit")
		next.ServeHTTP(writer, request)
	})
}

// BasicAuth - a handy middleware function that will provide basic auth around specific endpoints
/*
func BasicAuth(originalRequest func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		logLib.Info("basic auth endpoint hit")
		user , pass , ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			originalRequest(w,r)
		}else {
			w.Header().Set("Content-Type","application/json; charset=UTF-8")
			sendErrorResponse(w,"not authorized",errors.New("not authorized"))
			return
		}
	}
}
*/

// validateToken -
func validateToken(accessToken string) bool {
	var mySigningKey = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

// JwtAuth - a handy middleware function that will provide Jwt auth around specific endpoints
func JwtAuth(originalRequest func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logLib.Info("JWT auth endpoint hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		// Barer jwt-token
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		if validateToken(authHeaderParts[1]) {
			originalRequest(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}
	}
}
