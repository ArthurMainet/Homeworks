package main

import (
	"Email-API/config"
	"Email-API/internal"
	"log"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	repo := internal.NewLocalRepo()

	router := http.NewServeMux()
	info := internal.NewEmailHandler(internal.EmailHandlerDeps{
		Config: conf,
		Repo:   repo,
	})

	router.HandleFunc("POST /send", info.ReciveEmail())
	router.HandleFunc("GET /verify/{hash}", info.Verify())

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Starting to listen on port :8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
