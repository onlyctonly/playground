package main

import "fmt"

func main() {
	s:=make([]string,2)
	fmt.Println(len(s))

	s[0]=`a`
	s[1]=`b`
	fmt.Println(s)

	s = append(s, `c`, `e`)
	fmt.Println(s)
	fmt.Println(len(s))
}
