package main

import "fmt"

func main() {
	i := 1
	zeroval(i)
	fmt.Println(i)

	zeropoint(&i)
	fmt.Println(i)

}

func zeroval(v int) {
	v = 0
}
func zeropoint(p *int) {
	*p = 0
}
