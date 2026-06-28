package main

import(
	"net/http"
	"log"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	mux := http.NewServeMux()
	server := &http.Server {
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Server is starting on http://localhost:8080...")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}