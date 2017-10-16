package main

import (
	"log"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

	if req.Method != "POST" {
		var data []FoodItem
		var row FoodItem

		rows, err := db.Query("select id, name from food")
		handleError(err)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&row.Id, &row.Name)
			data = append(data, row)
			handleError(err)
		}

		err = rows.Err()
		handleError(err)

		temp, _ := template.ParseFiles("index.html")
		temp.Execute(res, data)

		return
	}
	
	formInput := req.FormValue("userInput")
	log.Println("form input: ", formInput)

	query, err := db.Prepare("INSERT food SET name=?")
	handleError(err)

	result, err := query.Exec(formInput)
	handleError(err)

	id, err := result.LastInsertId()
	handleError(err)

	log.Println("new entry in food row id#", id)
	
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
