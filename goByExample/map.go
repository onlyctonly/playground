package main

import "fmt"

func main() {
	m := make(map[string]int)
	m[`key1`] = 1
	m[`key2`] = 2
	fmt.Println(m[`key1`])
	v, ok := m[`key1`]
	fmt.Println(v, ok)
	delete(m, `key1`)
	fmt.Println(m)
	_, o := m[`key1`]
	fmt.Println(o)
}
