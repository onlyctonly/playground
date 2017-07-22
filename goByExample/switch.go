package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		switch i {
		case 1:
			fmt.Println(`one`)
		case 2:
			fmt.Println(`two`)
		default:
			fmt.Println(`the rest`)
		}
	}
}
