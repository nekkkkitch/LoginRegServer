package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	var users = GetUsers()
	tmpl := template.Must(template.ParseFiles("templates/MainPage.html"))
	tmpl.Execute(w, users)
}
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/LoginPage.html"))
	user := User{
		ID:       0,
		Username: r.FormValue("Username"),
		Password: r.FormValue("Password"),
	}
	tmpl.Execute(w, user)
	fmt.Println(user)
	//InsertUser(&user)
}
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
func InsertUser(u *User) {
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
func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/login/", Login)
	http.ListenAndServe("localhost:8080", nil)
}
