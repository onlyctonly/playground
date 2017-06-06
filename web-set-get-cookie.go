package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visited")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{Name: "visited", Value: "1"})
	} else {
		n,_:=strconv.Atoi(c.Value)
		n++
		http.SetCookie(w, &http.Cookie{Name: "visited", Value: strconv.Itoa(n) })
	}

}
func readCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-name")
	if err != nil {
		http.Error(w, "read cookie error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, c)

}
