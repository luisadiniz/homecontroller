package main

import (
	"fmt"
	"net/http"

	"github.com/luisadiniz/homecontroller/handlers"
	"github.com/luisadiniz/homecontroller/repositories"
)

func main() {
	//repo := repositories.NewInMemoryDB()
	repo, err := repositories.NewRelationalRepository()
	if err != nil {
		fmt.Println(err.Error())
	}
	router := http.NewServeMux()

	router.HandleFunc("/lightbulbs", handlers.HandleLightbulbs(repo))

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	fmt.Println("Server Running")

	server.ListenAndServe()
}
