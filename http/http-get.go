package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, _ := http.Get("http://www.apergia.gr/")
	defer resp.Body.Close()

	c, _:=ioutil.ReadAll(resp.Body)
	err:=ioutil.WriteFile("/Users/jxyu/apegia.txt", c, 0700)
	if err!=nil {
		fmt.Println(err)
	}


}
