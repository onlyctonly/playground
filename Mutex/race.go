package main

import (
	"sync"
	"fmt"
)
var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementer("foo: ")
	go incrementer("bar: ")
	wg.Wait()
	fmt.Println("final is: ", counter)
}


func incrementer(s string)  {
	for i:=0;i<20;i++ {
		x:=counter
		x++
		counter = x
		fmt.Println(s, i, "counter: ", counter)
	}
	wg.Done()
}