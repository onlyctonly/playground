package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	li, err:=net.Listen(`tcp`, ":8080")
	defer li.Close()
	check(err)
	for {
		conn, err:=li.Accept()
		check(err)
		go handle(conn)
	}
}

func check(e error) {
	if e!=nil {
		fmt.Println(e)
		return
	}
}
func handle(c net.Conn) {
	scanner:=bufio.NewScanner(c)
	for scanner.Scan() {
		ln:=scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(c,"you said: %s\n", ln)
	}
	defer c.Close()
}
