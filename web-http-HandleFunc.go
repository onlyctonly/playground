package main

import (
	"net/http"
	"fmt"
)

func cat(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "meowmeow")
}

func dog(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "wofwof")
}

func main()  {
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/dog", dog)

	http.ListenAndServe(":8080", nil)
}
