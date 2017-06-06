package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s:=`jxyu`
	e:=base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(e)
	ds,_:=base64.StdEncoding.DecodeString(e)
	fmt.Println(string(ds))
}
