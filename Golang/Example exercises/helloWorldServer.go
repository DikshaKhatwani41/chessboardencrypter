package main

import (
	"fmt"
	"log"
	"net/http"
)
// make a request to the server

// we get a request (we are server) then we print a response ("Hello, world")
// to the browser by the ResponseWriter (w) 
func homePage(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprint(w,"Hello, world")
}

func main(){
	http.HandleFunc("/",homePage)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080",nil)
}


// cd 1 - helloWorldServer
