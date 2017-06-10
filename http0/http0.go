package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)
type People struct {
	Name string
	Age int
}


func returnJson(w http.ResponseWriter, r *http.Request)  {
	p:=People{"xiangyu", 33}
	data,_:=json.Marshal(p)
	if r.Method == http.MethodGet {
		c:=http.Cookie{
			Name:"sid",
			Value:"123-456",
			HttpOnly:true,
		}
		http.SetCookie(w, &c)
		w.Header().Set("Content-Type", "text/json")
		//fmt.Println(r.Header.Get("Accept-Language"))
		fmt.Fprint(w,string(data))
	}
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", returnJson)
	http.HandleFunc("/showvalue", getCookie)
	http.ListenAndServe(":8080", nil)
}

func getCookie(w http.ResponseWriter, r *http.Request)  {
	c,err:=r.Cookie("sid")
	if err!=nil{
		fmt.Fprintln(w,"err")
		return
	}
	fmt.Fprintln(w,c.Value)
}
