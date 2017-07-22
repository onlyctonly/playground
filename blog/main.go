package main

import (
	"html/template"
	"net/http"
	//"github.com/satori/go.uuid"
        //"golang.org/x/crypto/bcrypt"
)
type user struct {
	Username string
	Password string
}
var tpl *template.Template
var dbusers=map[string]user{}
var dbsessions=map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	//check if logged in
	tpl.ExecuteTemplate(w,"index.html",nil)
}