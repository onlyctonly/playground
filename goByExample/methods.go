package main

import "fmt"

func main() {
	a := rect{5, 10}
	a.change()
	fmt.Println(a)
	a.notchange()
	fmt.Println(a)
}

type rect struct {
	width, height int
}

func (r *rect) change() {
	r.width = 10
}

func (r rect) notchange() {
	r.width = 0
}
