package main

import (
	"fmt"
	"log"
	"net/http"
)

//we get a request.(we are server). then we writing a response("Hello, world")
//to the browser by the ResponseWriter (w)
func homePage(w http.ResponseWriter, r *http.Request) {
	//Fprint writes into the w(ResponseWriter) instead of writting to the console
	//like fmt.print ! (so easy !!)
	fmt.Fprint(w, "Hello, world")
}

func main() {
	//we call this handler. "/" is the root or the address of our homepage
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	//this single line create a production server ready to serv thousands of requests!!!
	http.ListenAndServe(":8080", nil)
}
