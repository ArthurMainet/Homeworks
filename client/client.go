package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	client := http.Client{}

	resp, err := client.Get("http://localhost:8080/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status code: %d", resp.StatusCode)
}
