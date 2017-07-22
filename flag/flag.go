package main

import (
	"flag"
	"fmt"
)

func main() {
	p1 := flag.Int("ip1", 1234, "ip 1234")
	p2 := flag.Int("ip2", 8888, "ip 8888")
	flag.Parse()
	fmt.Println("ip1 value: ", *p1)
	fmt.Println("ip2 value: ", *p2)
	println(`what`) //std err; can be directed to log file: 2> log.txt
}
