package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{`z`, `a`, `x`}
	sort.Strings(a)
	fmt.Println(a)

	nums := []int{5, 899, 999, 1}
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println(sort.IntsAreSorted(nums))
}
