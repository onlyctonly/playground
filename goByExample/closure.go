package main

import "fmt"

func main() {
	nextint := intSeq()
	newint:=intSeq()
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(newint())
}

func intSeq() func() int  {
	i:=0
	return func() int {
		i++
		return i
	}
}
