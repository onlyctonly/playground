package main

import (
	"net/http"
	//"io"
	//"os"
	//"log"
)

func p(w http.ResponseWriter, r *http.Request)  {
	// solution 1
	//f, err:=os.Open("dog.jpg")
	//if err!=nil {
	//	log.Fatalln(err)
	//}
	//defer f.Close()
	//io.Copy(w, f)
	http.ServeFile(w,r,"dog.jpg")
}
func main()  {
	http.HandleFunc("/", p)
	http.ListenAndServe(":8080", nil)
}
