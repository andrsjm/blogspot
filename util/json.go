package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Count   int         `json:"count,omitempty"`
}

type jsonPresenter struct{}

type IPresenterJSON interface {
	SendError(w http.ResponseWriter, message string)
	SendSuccess(w http.ResponseWriter, data ...interface{})
}

func NewJsonPresenter() IPresenterJSON {
	return &jsonPresenter{}
}

func (*jsonPresenter) SendError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := BaseResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

func (*jsonPresenter) SendSuccess(w http.ResponseWriter, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")

	var response = BaseResponse{
		Status:  http.StatusOK,
		Message: "Success.",
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
