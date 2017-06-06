package main

import "fmt"

func main() {
	nums:=[]int{2,3,5,6,8,7}
	r := sums(nums...)
	fmt.Println(r)

}

func sums(nums ...int) int {
	var total int
	for _,v:=range nums {
		total+=v
	}
	return total
}