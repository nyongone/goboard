package handler

import (
	"encoding/json"
	"go-board-api/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	Service model.CommentService
}

func (ch *CommentHandler) GetCommentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	comment, err := ch.Service.FindOne(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusNotFound,
			Message: "NOT_FOUND",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Response{
		Code: http.StatusOK,
		Message: "OK",
		Data: comment,
	})
}

func (ch *CommentHandler) GetCommentsByPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	comments, err := ch.Service.FindAllByPost(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Response{
		Code: http.StatusOK,
		Message: "OK",
		Data: comments,
	})
}

func (ch *CommentHandler) GetCommentByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	comments, err := ch.Service.FindAllByAuthor(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Response{
		Code: http.StatusOK,
		Message: "OK",
		Data: comments,
	})
}

func (ch *CommentHandler) WriteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var comment model.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})

		return
	}

	authUser := r.Context().Value(model.UserKey).(*model.User)
	comment.AuthorID = authUser.ID

	err := ch.Service.Create(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&model.Response{
		Code: http.StatusCreated,
		Message: "CREATED",
	})
}

func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := ch.Service.Delete(id, true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})

		return
	}

	authUser := r.Context().Value(model.UserKey).(*model.User)
	c, err := ch.Service.FindOne(id)
	if err != nil || c.Author.ID != authUser.ID {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&model.Response{
			Code: http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Response{
		Code: http.StatusOK,
		Message: "OK",
	})
}