package warehouse

import "time"

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

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		message: err.Error(),
		time:    time.Now(),
	}
}

func NewItem(name string, quantity int, minTh int) Item {
	return Item{
		Name:         name,
		Quantity:     quantity,
		MinThreshold: minTh,
	}
}
