package http

import (
	"encoding/json"
	"github.com/Kratos40-sba/go-prod/internal/comment"
	"github.com/gorilla/mux"
	logLib "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

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
		logLib.Error(err)
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
		logLib.Error(err)
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
		logLib.Error(err)
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
		logLib.Error(err)
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
		logLib.Error(err)
	}
}
