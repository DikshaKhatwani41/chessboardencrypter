package main

import (
	"fmt"
	"log"
	"net/http"
)

// make a request to the server

// we get a request (we are server) then we print a response ("Hello, world")
// to the browser by the ResponseWriter (w)
func homePage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello world</strong>`
	w.Header().Set("Content-Type", "text/html")

	//fmt.Fprint(w, "Hello, world")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

// cd 1 - helloWorldServer
