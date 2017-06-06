package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main()  {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Document</title>
		</head>
		<body>
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="f">
		<input type="submit">
		</body>
		</html>
	`)
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("f")
		if err!=nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		//fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, "read file err", http.StatusInternalServerError)
			return
		}

		df, err := os.Create(filepath.Join("./templates/", h.Filename))
		if err!=nil {
			http.Error(w, "create file error", http.StatusInternalServerError)
			return
		}
		defer df.Close()
		_, err = df.Write(bs)
		if err!=nil {
			http.Error(w, "write file error", http.StatusInternalServerError)
			return
		}

	}
}
