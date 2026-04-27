package warehouse

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Item struct {
	Name         string
	Quantity     int
	MinThreshold int
}

type ItemDTO struct {
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	MinThreshold int    `json:"minThreshold"`
}
type ErrorDTO struct {
	message string
	time    time.Time
}

type ItemQuantityDTO struct {
	Quantity int `json:"quantity"`
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		message: err.Error(),
		time:    time.Now(),
	}
}

func IsError(w http.ResponseWriter, err error) {
	errDTO := NewErrorDTO(err)
	if errors.Is(err, ErrItemAlreadyExist) {
		http.Error(w, errDTO.ToString(), http.StatusConflict)
	} else if errors.Is(err, ErrItemNotFound) {
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
	}
}

func NewItem(name string, quantity int, minTh int) Item {
	return Item{
		Name:         name,
		Quantity:     quantity,
		MinThreshold: minTh,
	}
}
