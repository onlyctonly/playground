package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	//open file
	f,err:=os.Open("lineFilter/test.txt")
	check(err)
	defer f.Close()
	//scan file
	scan:=bufio.NewScanner(f)
	//create new file
	f2,err:=os.Create("lineFilter/test2.txt")
	check(err)
	defer f2.Close()
	//remove first line & create new file
	for n:=0;scan.Scan();n++ {
		if n>0 {
			_,err:=f2.WriteString(scan.Text()+"\n")
			check(err)
		}
	}
}
func check(e error) {
	if e!=nil {
		fmt.Println(e)
	}
}