package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func RandomHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand := rand.IntN(6) + 1
		randStr := strconv.Itoa(rand)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(randStr))
	}
}

func main() {
	//creating router
	//sadkaskjdjas
	//aslkdmklasmdklasd
	//askdmkasdmklaskldm
	router := http.NewServeMux()
	router.HandleFunc("/random", RandomHandler())

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
