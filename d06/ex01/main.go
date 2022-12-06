package main

import (
	"fmt"
	"github.com/go-pg/pg"
	// "github.com/go-pg/pg/orm"
	"html/template"
	"net/http"
	"os"
	// "github.com/go-pg/pg"
	// "github.com/go-pg/pg/orm"
)

type Post struct {
	Id      int64
	Title   string
	Content string
	// AuthorEmail string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/homePage.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := DBConn()
	var posts []Post

	err = db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	// posts[0].Content = "first"
	// posts[1].Content = "second"
	extra := struct {
		Posts []Post
	}{Posts: posts}

	tmpl.Execute(w, extra)
}

func DBConn() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		Database: "blog_db",
		User:     "blog",
		Password: "blog_secret_password",
	})
	return db
}

func addContentInDB() {
	db := DBConn()
	var posts []Post

	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	if posts[0].Content == "" {
		posts[0].Content = "first"
	}
	if posts[0].Content == "" {
		posts[1].Content = "second"
	}

}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
