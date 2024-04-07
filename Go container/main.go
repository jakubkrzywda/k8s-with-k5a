package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	fmt.Println("Server is running at 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe( "0.0.0.0:8080", mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
	log.Print("hello")
}
