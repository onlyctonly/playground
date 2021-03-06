package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementer("Foo: ")
	go incrementer("Bar: ")
	wg.Wait()
	fmt.Println("final counter: ", counter)
}

func incrementer(s string) {
	for i := 0; i < 20; i++ {
		counter++
		fmt.Println(s, i, "counter: ", counter)
	}
	wg.Done()
}
