package main

import (
	"log"
	"net/http"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandle(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("static/view.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandle)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
