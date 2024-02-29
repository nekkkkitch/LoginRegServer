package main

import "database/sql"

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
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
