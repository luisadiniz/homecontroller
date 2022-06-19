package main

import (
	"fmt"
	"net/http"

	"github.com/luisadiniz/homecontroller/handlers"
	"github.com/luisadiniz/homecontroller/repositories"
)

func main() {
	repo := repositories.New()
	router := http.NewServeMux()

	router.HandleFunc("/lightbulbs", handlers.HandleLightbulbs(repo))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server Running")

	server.ListenAndServe()
}
