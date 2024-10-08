package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func viewHandle(w http.ResponseWriter, r *http.Request) {
	signatures := getStrings("signature.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("static/view.html")
	check(err)
	guestbook := GuestBook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(w, guestbook)
	check(err)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("static/new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signature.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandle)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
