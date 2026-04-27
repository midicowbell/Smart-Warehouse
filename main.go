package main

import (
	"fmt"
	"warehouse/http"
	"warehouse/warehouse"
)

func main() {
	repo := warehouse.NewRepository()
	HTTPHandlers := http.NewHTTPHandlers(repo)
	server := http.NewHTTPServer(HTTPHandlers)
	if err := server.StartServer(); err != nil {
		fmt.Println("failed to start HTTP server: ", err)
	}
}
