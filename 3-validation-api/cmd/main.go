package main

import (
	"Email-API/packages"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	info := packages.NewEmailHandler()

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
