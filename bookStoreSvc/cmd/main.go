package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thtrangphu/bookStoreSvc/config"
	"github.com/thtrangphu/bookStoreSvc/entity"
	"log"
	"net/http"
)

func main() {
	// Configure the database connection (always check errors)
	//db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DATABASE_USERNAME, config.DATABASE_PASSWORD, config.DATABASE_HOST, config.DATABASE_DBNAME))

	if err != nil {
		log.Panic(err.Error())
	}

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println(entity.CreateNewUser("hello", "hello", "PTKTrang"))
	http.ListenAndServe(config.PORT, nil)
}
