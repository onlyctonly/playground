package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	for index, v := range nums {
		fmt.Println(index, ": ", v)
	}
}
