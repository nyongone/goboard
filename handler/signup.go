package handler

import (
	"encoding/json"
	"go-board-api/model"
	"net/http"
)

type SignupHandler struct {
  Service model.SignupService
}

func (sh *SignupHandler) Signup(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var signup model.Signup
  if err := json.NewDecoder(r.Body).Decode(&signup); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusBadRequest,
      Message: "BAD_REQUEST",
    })

    return
  }

  err := sh.Service.Signup(&signup)
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