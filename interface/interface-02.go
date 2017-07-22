package main

import "fmt"

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(shapes ...shape) {
	for _, v := range shapes {
		fmt.Println("this is ", v)
		fmt.Println(v.area())
	}
}

func main() {
	s := square{8}
	q := square{9}
	info(s, q)

}

// ...shape
