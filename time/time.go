package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	d := time.Now().Add(time.Hour * 24).Day()
	m := int(time.Now().Add(time.Hour * 24).Month())
	d2 := "0" + strconv.Itoa(d)
	fmt.Println(d2)
	fmt.Println(m)
}
