package main

import (
	"time"
	"fmt"
)

func main() {
	t1:=time.Date(2017,07,23,03,17,27,0,time.UTC)
	t2:=t1.Add(-time.Millisecond*171804143)
	fmt.Println(t2)
}
