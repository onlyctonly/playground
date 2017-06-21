package main

import "fmt"

func main() {
	c:=make(chan int)

	go func() {
		for i:=0;i<10;i++ {
			c<-i
		}
		close(c) // if no close, there will be deadlock
	}()

	for v:=range c {
		fmt.Println(v)
	}
}
