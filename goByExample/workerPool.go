package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= 4; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 4; a++ {
		<-result
	}
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println(`worder `, id, `started job`, j)
		time.Sleep(time.Second * 1)
		fmt.Println(`worker `, id, `finished job`, j)
		result <- j
	}
}
