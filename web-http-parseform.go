package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int
var t *template.Template


func (handler hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		f1 map[string][]string
		f2 map[string][]string
	}{
		req.Form,
		req.PostForm,
	}

	t.ExecuteTemplate(w, "web-http-parseform.gohtml", data)
}


func init() {
	t = template.Must(template.ParseFiles("web-http-parseform.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)

}
