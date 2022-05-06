package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhosst:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
