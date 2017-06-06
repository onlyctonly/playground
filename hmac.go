package main

import (
	"crypto/hmac"
	"crypto/sha256"
	//"io"
	"fmt"
	"io"
)

func main() {
	h1:=hmac.New(sha256.New, []byte(`key`))
	h2:=hmac.New(sha256.New, []byte(`key`))
	s:=`jxyu`
	h1.Write([]byte(s))
	fmt.Printf("%x\n", h1.Sum(nil) )
	io.WriteString(h2, "jxyu")
	fmt.Printf("%x\n", h2.Sum(nil))

	fmt.Println(hmac.Equal(h1.Sum(nil), h2.Sum(nil)))
}
