package main

import (
	"Email-API/config"
	"Email-API/internal/products"
	"Email-API/internal/verify"
	"Email-API/packages/db"
	"Email-API/packages/middlewares"
	"fmt"
	"log"
	"net/http"
)

// Added new information to check HW

func main() {
	conf := config.LoadConfig()
	db := db.NewDB(conf)
	localrepo := verify.NewLocalRepo()
	productRepo := products.NewProductRepository(db)
	fmt.Println(productRepo)

	router := http.NewServeMux()
	verify.NewEmailHandler(router, verify.EmailHandlerDeps{
		Config: conf.EmailConf,
		Repo:   localrepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middlewares.Logging(router),
	}

	log.Println("Starting to listen on port :8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
