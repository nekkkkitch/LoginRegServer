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
	fmt.Println(users)
	tmpl := template.Must(template.ParseFiles("templates/MainPage.html"))
	tmpl.Execute(w, users)
}
func GetUsers() []User {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")
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
func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", MainPage)
	http.ListenAndServe("localhost:8080", nil)
}
