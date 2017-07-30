package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)
var db *sql.DB
type Book struct {
	Id int
	Name string
	Author string
	Isbn string
}
func init() {
	var err error
	if db, err = sql.Open("mysql", "root:root@/store");err!=nil {
		log.Fatalln("db failure")
	}
	if err=db.Ping();err!=nil {
		panic("cannot connet to db")
	}
	fmt.Println("db is connected")
}

func main() {
	http.HandleFunc("/books", books)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func books(w http.ResponseWriter, r *http.Request) {
	if r.Method!=http.MethodGet {
		http.Error(w,http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err:=db.Query("select * from books;")
	if err!=nil {
		fmt.Println(err)
	}

	defer rows.Close()
	bks:=[]Book{}
	for rows.Next() {
		bk:=Book{}
		err:=rows.Scan(&bk.Id,&bk.Name,&bk.Author,&bk.Isbn)
		if err!=nil {
			http.Error(w,http.StatusText(500), http.StatusInternalServerError)
			return
		}
		bks=append(bks,bk)
	}
	if err:=rows.Err();err!=nil{
		http.Error(w,http.StatusText(500), http.StatusInternalServerError)
		return
	}
	for _,bk:=range bks {
		fmt.Fprintf(w, "%d, %s, %s, %s\n", bk.Id, bk.Name,bk.Author,bk.Isbn)
	}

}
