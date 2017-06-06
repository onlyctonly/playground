package main

import "fmt"

func main() {
	i:=1
	for i<=10 {
		fmt.Println(`first: `, i)
		i++
	}

	for j:=0;j<=10;j++ {
		if j%2 != 0 {
			continue
		}
		fmt.Println(`sencond:`,j)
	}
}
