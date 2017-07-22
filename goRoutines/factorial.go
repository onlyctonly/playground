package main

import "fmt"

func main() {
	c := make(chan int)
	fac := 1
	go func() {
		for n := 7; n > 0; n-- {
			fac *= n
		}
		c <- fac
	}()
	fmt.Println(<-c)
}
