package handler

import (
	"encoding/json"
	"go-board-api/model"
	"net/http"
)

type BoardHandler struct {
	Service model.BoardService
}

func (bh *BoardHandler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	boards, err := bh.Service.FindAll()
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
		Data: boards,
	})
}