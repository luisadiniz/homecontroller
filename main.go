package main

import (
	"fmt"
	"net/http"

	"github.com/luisadiniz/homecontroller/handlers"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/lightbulbs", handlers.HandleLightbulbs)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server Running")

	server.ListenAndServe()
}
