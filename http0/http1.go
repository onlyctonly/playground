package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request)  {

	if r.Method == http.MethodGet {
		c,err:=r.Cookie("sid")
		if err == http.ErrNoCookie {
			tpl.ExecuteTemplate(w,"login.gohtml", nil)
			return
		} else if err !=nil {
			http.Error(w,err,http.StatusInternalServerError)
			return
		} else {
			//compare sid with db
				//if ok, redirect

				//if not ok, go to login
		}
	}


	if r.Method == http.MethodPost {

	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w,"register.gohtml", nil)
		return
	}
	if r.Method == http.MethodPost {
		//if username exits

		// not exits writes to db and redirect to content

		// exits
	}
}