package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
)

var (
	db *sql.DB 
	err error
)

type FoodItem struct {
	Id int
	Name string
}

func home(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:sdfccsdfcc@(127.0.0.1:3306)/sphere")
	handleError(err)
	defer db.Close()

	var data []FoodItem

	if req.Method != "POST" {
		var (
			id int
			name string
		)

		rows, err := db.Query("select id, name from food")
		handleError(err)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id, &name)
			data = append(data, FoodItem{id, name})
			handleError(err)
		}

		err = rows.Err()
		handleError(err)

		temp, _ := template.ParseFiles("index.html")
		temp.Execute(res, data)

		return
	}
	
	formInput := req.FormValue("userInput")
	log.Println(formInput)
	
	http.Redirect(res, req, "/", 301)

}

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":4200", nil)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

