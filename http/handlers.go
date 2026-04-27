package http

import (
	"encoding/json"
	"net/http"
	"warehouse/warehouse"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item warehouse.ItemDTO
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		err := warehouse.NewErrorDTO(err)
		json.MarshalIndent(err, "", "    ")
		w.WriteHeader(http.StatusBadRequest)
	}
}
