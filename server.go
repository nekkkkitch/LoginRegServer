package main

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("bcd.html", "main.css")
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", MainPage)
	http.ListenAndServe("localhost:8080", nil)
}
