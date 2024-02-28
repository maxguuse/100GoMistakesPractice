package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

var DB *sql.DB

// Bad init function, won't allow user to handle errors
func init() {
	dsn := os.Getenv("DSN")
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}

	DB = d
}

// Good init function, won't panic, just setup the default HTTP handlers
func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)
	http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc/", static))
}

func main() {

}
