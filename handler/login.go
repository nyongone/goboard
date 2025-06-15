package handler

import (
	"encoding/json"
	"go-board-api/internal/util"
	"go-board-api/model"
	"net/http"
)

type LoginHandler struct {
  Service	model.LoginService
}

func (lh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var login model.Login
  if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusBadRequest,
      Message: "BAD_REQUEST",
    })

    return
  }

  user, err := lh.Service.FindOneByEmail(login.Email)
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusUnauthorized,
      Message: "UNAUTHORIZED",
    })

    return
  }

  if isAuthorized, err := util.ComparePassword(user.Password, login.Password); !isAuthorized || err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusUnauthorized,
      Message: "UNAUTHORIZED",
    })

    return
  }

  access_token, err := lh.Service.CreateAccessToken(user.Email)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusInternalServerError,
      Message: "INTERNAL_SERVER_ERROR",
    })

    return
  }

  refresh_token, err := lh.Service.CreateRefreshToken()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusInternalServerError,
      Message: "INTERNAL_SERVER_ERROR",
    })

    return
  }

  err = lh.Service.UpdateRefreshToken(user.ID, refresh_token)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusInternalServerError,
      Message: "INTERNAL_SERVER_ERROR",
    })

    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(&model.Response{
    Code: http.StatusOK,
    Message: "OK",
    Data: &model.JWT{
      AccessToken: access_token,
      RefreshToken: refresh_token,
    },
  })
}