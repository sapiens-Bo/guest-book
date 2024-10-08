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

func main() {
	http.HandleFunc("/guestbook", viewHandle)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
