package handler

import (
	"encoding/json"
	"go-board-api/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
  Service model.UserService
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  user, err := uh.Service.FindOne(id)
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
    Data: user,
  })
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  
  var user model.User
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusBadRequest,
      Message: "BAD_REQUEST",
    })

    return
  }

  authUser := r.Context().Value(model.UserKey).(*model.User)
  if id != authUser.ID {
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusUnauthorized,
      Message: "UNAUTHORIZED",
    })

    return
  }

  err := uh.Service.Update(id, &user)
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