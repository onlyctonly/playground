package main

import "fmt"

func main() {
	x := 5
	y := &x
	fmt.Println(x)
	fmt.Println(&x)

	*y=7
	fmt.Println(y)
	fmt.Println(x)

}
