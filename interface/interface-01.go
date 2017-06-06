package main

import "fmt"

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side*s.side
}
type shape interface {
	area() float64
}

func info(s shape)  {
	fmt.Println("this is ", s)
	fmt.Println(s.area())
}
func main() {
	s:=square{8.0}
	info(s)

}
