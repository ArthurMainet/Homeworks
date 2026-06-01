package main

import (
	"Email-API/config"
	"Email-API/internal/verify"
	"log"
	"net/http"
)

// Added new information to check HW

func main() {
	conf := config.LoadConfig()
	repo := verify.NewLocalRepo()

	router := http.NewServeMux()
	info := verify.NewEmailHandler(verify.EmailHandlerDeps{
		Config: conf.EmailConf,
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
