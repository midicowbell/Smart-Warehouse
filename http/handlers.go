package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"warehouse/warehouse"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	repo *warehouse.Repository
}

func NewHTTPHandlers(repos *warehouse.Repository) *HTTPHandlers {
	return &HTTPHandlers{
		repo: repos,
	}
}

func (h *HTTPHandlers) HandleCreateItem(w http.ResponseWriter, r *http.Request) {
	var itemDTO warehouse.ItemDTO
	if err := json.NewDecoder(r.Body).Decode(&itemDTO); err != nil {
		errDTO := warehouse.NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	item := warehouse.NewItem(itemDTO.Name, itemDTO.Quantity, itemDTO.MinThreshold)
	if err := h.repo.AddItem(item); err != nil {
		warehouse.IsError(w, err)
		return
	}
	b, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response HTTP: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	item, err := h.repo.GetItem(name)
	if err != nil {
		warehouse.IsError(w, err)
		return
	}
	b, err := json.MarshalIndent(item, "", "   ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed response to HTTP: ", err)
	}
}

func (h *HTTPHandlers) HandleListItems(w http.ResponseWriter, r *http.Request) {
	ans := h.repo.ListItems()
	b, err := json.MarshalIndent(ans, "", "   ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response HTTP: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleListLowStockItems(w http.ResponseWriter, r *http.Request) {
	ans := h.repo.ListLowStockItems()
	b, err := json.MarshalIndent(ans, "", "   ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response HTTP: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleUpdateQuantity(w http.ResponseWriter, r *http.Request) {
	var itemQuantityDTO warehouse.ItemQuantityDTO
	if err := json.NewDecoder(r.Body).Decode(&itemQuantityDTO); err != nil {
		errDTO := warehouse.NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	name := mux.Vars(r)["name"]
	if err := h.repo.UpdateQuantity(name, itemQuantityDTO.Quantity); err != nil {
		warehouse.IsError(w, err)
		return
	}
	item, err := h.repo.GetItem(name)
	if err != nil {
		warehouse.IsError(w, err)
		return
	}
	b, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response HTTP: ", err)
		return
	}

}

func (h *HTTPHandlers) HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := h.repo.DeleteItem(name); err != nil {
		warehouse.IsError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
