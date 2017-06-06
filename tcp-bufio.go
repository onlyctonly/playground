package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()
	if err !=nil {
		log.Fatal(err)
	}

	for {
		conn, err := li.Accept()

		if err !=nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(c net.Conn)  {
	scanner:=bufio.NewScanner(c)
	for scanner.Scan() {
		line:=scanner.Text()
		fmt.Println(line)
	}
	defer c.Close()
}
