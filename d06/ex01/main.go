package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/gorilla/context"
	// "golang.org/x/time/rate"
	// "github.com/go-pg/pg/orm"
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// var limiter = rate.NewLimiter(1, 3)

type Post struct {
	Id      int64
	Title   string
	Content string
	// AuthorEmail string
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/adminPage.html", "templates/showAll.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := DBConn()
	var posts []Post
	defer db.Close()
	err = db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	extra := struct {
		Posts []Post
	}{Posts: posts}
	tmpl.Execute(w, extra)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/homePage.html", "templates/showAll.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := DBConn()
	var posts []Post
	defer db.Close()
	err = db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	var countPosts int

	for countPosts = range posts {
	}

	// fmt.Println("countPosts = ", countPosts)
	countPosts++
	pages := countPosts/3 + 1
	// fmt.Println("pages = ", pages)
	// nId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	// if nId > 0{

	// }
	var posts1 []Post
	var pg []int = make([]int, pages)

	for i := 0; i < pages; i++ {
		pg[i] = i
	}
	nId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		nId = 0
	}
	// fmt.Println("nId = ", nId)
	// fmt.Println("len = ", len(posts))
	if len(posts) > 0+(int(nId)*3) {
		posts1 = append(posts1, posts[0+(nId*3)])
	}
	if len(posts) > 1+(int(nId)*3) {
		posts1 = append(posts1, posts[1+(nId*3)])
	}
	if len(posts) > 2+(int(nId)*3) {
		posts1 = append(posts1, posts[2+(nId*3)])
	}
	// posts1 = append(posts1, posts[0+(nId*3)], posts[1+(nId*3)], posts[2+(nId*3)])
	extra := struct {
		Posts []Post
		Pages []int
	}{Posts: posts1, Pages: pg}
	// _ = pg
	// extra := struct {
	// 	Posts []Post
	// }{Posts: posts1}
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

func DeleteContentInDB() {
	db := DBConn()
	var posts []Post
	defer db.Close()
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	if posts != nil {
		res, err := db.Model(&posts).Delete()
		if err != nil {
			panic(err)
		}
		fmt.Println("deleted", res.RowsAffected())
	}
}

func addContentInDB() {
	db := DBConn()
	post := Post{
		Title:   "title 3",
		Content: "third content",
	}
	defer db.Close()
	err := db.Insert(&post)
	if err != nil {
		panic(err)
	}
}

func showPage(w http.ResponseWriter, r *http.Request) {
	nId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	// fmt.Fprintf(w, "nid =  %d", nId)

	tmpl, err := template.ParseFiles("templates/homePage.html", "templates/showOne.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := DBConn()
	var posts []Post
	defer db.Close()
	err = db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	var i int
	for i = range posts {
		if posts[i].Id == nId {
			break
		}
	}

	fmt.Println("i =", i)
	fmt.Println("posts[i].Id =", posts[i].Id)

	extra := struct {
		Id      int64
		Title   string
		Content string
	}{Id: posts[i].Id, Title: posts[i].Title, Content: posts[i].Content}

	tmpl.Execute(w, extra)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	data := context.Get(r, "data")
	t, _ := template.ParseFiles("templates/homePage.html", "templates/login.html")
	t.Execute(w, data)
}

func loginCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rrrrr")
	if r.Method == "POST" {
		login := r.FormValue("login")
		passw := r.FormValue("passw")

		file, err := os.Open("templates/admin_credentials.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		fileScanner.Scan()
		fmt.Println(fileScanner.Text())
		correctLog := fileScanner.Text()
		fileScanner.Scan()
		fmt.Println(fileScanner.Text())
		correctPassw := fileScanner.Text()

		if correctPassw == passw && correctLog == login {
			fmt.Println("coorect")
			http.Redirect(w, r, "/admin", 301)
		} else {
			fmt.Println("not coorect")
			http.Redirect(w, r, "/", 301)
		}
	}
}

func newPostPage(w http.ResponseWriter, r *http.Request) {
	data := context.Get(r, "data")
	t, _ := template.ParseFiles("templates/homePage.html", "templates/new.html")
	t.Execute(w, data)
}

func newPostInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := DBConn()
		title := r.FormValue("title")
		content := r.FormValue("content")
		post1 := &Post{
			Title:   title,
			Content: content,
		}
		err := db.Insert(post1)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}

func main() {
	// DeleteContentInDB()
	// addContentInDB()

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/show", showPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/login/insert/", loginCheck)
	http.HandleFunc("/admin", adminPage)
	http.HandleFunc("/admin/newPost", newPostPage)
	http.HandleFunc("/admin/newPost/insert/", newPostInsert)
	http.ListenAndServe(":8888", nil)
}
