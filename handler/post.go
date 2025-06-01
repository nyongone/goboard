package handler

import (
	"encoding/json"
	"go-board-api/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
  Service model.PostService
}

func (ph *PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  post, err := ph.Service.FindOne(id)
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
    Data: post,
  })
}

func (ph *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  posts, err := ph.Service.FindAll()
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
    Data: posts,
  })
}

func (ph *PostHandler) GetAllPostsByBoard(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  board_id, _ := strconv.Atoi(vars["board_id"])
  posts, err := ph.Service.FindAllByBoard(board_id)
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
    Data: posts,
  })
}

func (ph *PostHandler) GetAllPostsByAuthor(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  author_id, _ := strconv.Atoi(vars["author_id"])
  posts, err := ph.Service.FindAllByAuthor(author_id)
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
    Data: posts,
  })
}

func (ph *PostHandler) WritePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var post model.Post
  if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusBadRequest,
      Message: "BAD_REQUEST",
    })

    return
  }

  authUser := r.Context().Value(model.UserKey).(*model.User)
  post.AuthorID = authUser.ID

  err := ph.Service.Create(&post)
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

func (ph *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  
  var post model.Post
  if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusBadRequest,
      Message: "BAD_REQUEST",
    })

    return
  }

  authUser := r.Context().Value(model.UserKey).(*model.User)
  p, err := ph.Service.FindOne(id)
  if err != nil || p.Author.ID != authUser.ID {
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusUnauthorized,
      Message: "UNAUTHORIZED",
    })

    return
  }

  err = ph.Service.Update(id, &post)
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
  })
}

func (ph *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])

  err := ph.Service.Delete(id, true)
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
  })
}