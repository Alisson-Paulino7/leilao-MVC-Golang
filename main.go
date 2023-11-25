package main

import (
	"database/sql"
	"goapp/routes"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "go_app_user:go_app_password@tcp(mysql:3306)/go_app_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
    //     http.ServeFile(w, r, "app/templates"+r.URL.Path)
    // })

	fs := http.FileServer(http.Dir("/app/templates"))
	http.Handle("/", fs)

	routes.CarregaRotas()

	http.ListenAndServe(":8080", nil)
}
