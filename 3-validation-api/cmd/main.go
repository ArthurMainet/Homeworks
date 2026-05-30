package main

import (
	"Email-API/config"
	"Email-API/packages"
	"log"
	"net/http"
)

func main() {
	conf := config.LoadConfig()

	router := http.NewServeMux()
	info := packages.NewEmailHandler(packages.EmailHandlerDeps{Config: conf})

	router.HandleFunc("POST /send", info.Send())
	router.HandleFunc("GET /verify/{hash}", info.Verify())

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
