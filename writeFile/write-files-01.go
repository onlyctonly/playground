package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	d1 := []byte("hello\ngo\n")
	err:=ioutil.WriteFile("/Users/jxyu/d1", d1, 0600)
	check(err)

	f, err:=os.Create("/Users/jxyu/d2")
	defer f.Close()
	n1, err:=f.Write(d1)
	check(err)
	fmt.Printf("wrote %d bytes\n", n1)

	n2, err:=f.WriteString("hello again")
	check(err)
	fmt.Printf("2nd wrote %d bytes\n", n2)

	f.Sync()




}

func check(e error)  {
	if e !=nil {
		panic(e)
	}
}