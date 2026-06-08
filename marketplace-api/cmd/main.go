package main

import (
	"Email-API/config"
	"Email-API/internal/auth"
	"Email-API/internal/products"
	"Email-API/internal/user"
	"Email-API/internal/verify"
	"Email-API/packages/db"
	"Email-API/packages/middlewares"
	"log"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	db := db.NewDB(conf)
	localrepo := verify.NewLocalRepo()
	productRepo := products.NewProductRepository(db)
	userRepo := user.NewUserRepository(db)

	router := http.NewServeMux()

	// Services
	emailService := verify.NewEmailService(&verify.EmailServiceDeps{
		Repo:           localrepo,
		EmailConf:      conf.EmailConf,
		UserRepository: userRepo,
	})
	phoneService := verify.NewPhoneService(&verify.PhoneServiceDeps{
		Repo:           localrepo,
		UserRepository: userRepo,
		JWT:            conf.AuthToken,
	})
	authService := auth.NewAuthService(&auth.AuthServiceDeps{
		Repo:         userRepo,
		EmailService: emailService,
		PhoneService: phoneService,
		JWT:          conf.AuthToken,
	})

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		EmailService: emailService,
		PhoneService: phoneService,
		Repo:         localrepo,
	})
	products.NewProductHandler(router, products.ProductHandlerDeps{
		ProductRepository: productRepo,
		Config:            conf.AuthToken,
	})
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		AuthService: authService,
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
