package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<a href="/set">set cookie</a>`)
}
func set(w http.ResponseWriter, r *http.Request)  {
	http.SetCookie(w, &http.Cookie{
		Name:"username",
		Value:"sadfasdfasdfasdf",
		MaxAge:10,
	})

	http.Redirect(w,r,"/read",http.StatusSeeOther)
}
func read(w http.ResponseWriter, r *http.Request) {
	c,err:=r.Cookie("username")
	if err!=nil {
		fmt.Fprint(w,err)
		return
	}
	fmt.Fprint(w,c.Value)
}