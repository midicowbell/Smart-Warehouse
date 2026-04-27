package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(httpHand *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: httpHand,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/items").Methods("POST").HandlerFunc(s.handlers.HandleCreateItem)
	router.Path("/items/{name}").Methods("GET").HandlerFunc(s.handlers.HandleGetItem)
	router.Path("/items").Methods("GET").Queries("status", "low").HandlerFunc(s.handlers.HandleListLowStockItems)
	router.Path("/items").Methods("GET").HandlerFunc(s.handlers.HandleListItems)
	router.Path("/items/{name}").Methods("PATCH").HandlerFunc(s.handlers.HandleUpdateQuantity)
	router.Path("/items/{name}").Methods("DELETE").HandlerFunc(s.handlers.HandleDeleteItem)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
