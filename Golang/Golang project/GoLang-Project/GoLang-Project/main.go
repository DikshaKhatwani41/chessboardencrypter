package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/chess"
	"net/http"
)

// Limitation
// Doesn't work for 6,7,8,9 numbers as pieces available were 32 in total
// alphabets took 26 pieces and we were left with 6 pieces only
// 6,7,8,9 are not acceptable

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func encode(w http.ResponseWriter, r *http.Request) {
	//i have added here playerChoice variable
	//from object request, we getting variable "str" content
	playerChoice := r.URL.Query().Get("str")
	result := chess.Encode(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
func decode(w http.ResponseWriter, r *http.Request) {

	playerChoice := r.URL.Query().Get("str")
	result := chess.Decode(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
