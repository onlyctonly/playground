package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 4 //waiting here
	fmt.Println(<-c)
}
