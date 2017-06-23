package main

import "fmt"

func main() {
	c:=make(chan int)
	c<-4
	fmt.Println(<-c)
}