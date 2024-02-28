package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/login/", LoginPage)
	http.HandleFunc("/login_user", LoginUser)
	http.ListenAndServe("localhost:8080", nil)
}

// http pages
func MainPage(w http.ResponseWriter, r *http.Request) {
	var users = GetUsers()
	tmpl := template.Must(template.ParseFiles("templates/MainPage.html"))
	tmpl.Execute(w, users)
}

var loginError string

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/LoginPage.html"))
	tmpl.Execute(w, loginError)
	loginError = "abc"
}

// http func
func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user = User{0, r.Form["Username"][0], r.Form["Password"][0]}
	if !CheckForSameLoginUser(user.Username) {
		InsertUser(user)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		loginError = "User with that login already exists :("
	}
}

// sql requests
func GetUsers() []User {
	db, err := sql.Open("mysql", "root:JonnekJuar4002@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	data, err := db.Query("select * from users")
	var result = []User{}
	for data.Next() {
		var user User
		err := data.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, user)
	}
	return result
}
func CheckForSameLoginUser(name string) bool {
	db, err := sql.Open("mysql", "root:JonnekJuar4002@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	row := db.QueryRow("select exists(select * from Users where Username = ?)", name)
	exist := false
	row.Scan(&exist)
	return exist
}
func InsertUser(u User) {
	db, err := sql.Open("mysql", "root:JonnekJuar4002@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("insert into users(Username, Password) values (?, ?)", u.Username, u.Password)
	if err != nil {
		panic(err.Error())
	}
}
