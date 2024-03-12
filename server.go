package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/login/", LoginPage)
	http.HandleFunc("/login_user", LoginUser)
	http.HandleFunc("/signup/", SignUpPage)
	http.HandleFunc("/signup_user", SignUpUser)
	http.ListenAndServe("localhost:8080", nil)
}

// http pages
var tmpl *template.Template

func MainPage(w http.ResponseWriter, r *http.Request) {
	var users = GetUsers()
	tmpl = template.Must(template.ParseFiles("templates/MainPage.html"))
	tmpl.Execute(w, users)
}

var loginError string

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/LoginPage.html"))
	tmpl.Execute(w, loginError)
	loginError = ""
}

var signUpError string

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/SignUpPage.html"))
	tmpl.Execute(w, signUpError)
	signUpError = ""
}

// http func
func SignUpUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user = User{0, r.Form["Login"][0], r.Form["Username"][0], r.Form["Password"][0]}
	if !CheckForSameLoginUser(user.Username) {
		InsertUser(user)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		signUpError = "User with that login already exists :("
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}
