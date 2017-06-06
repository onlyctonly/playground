package main

import (
	"net"
	"log"
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
		fmt.Fprintln(conn, "Hello World!")
		conn.Close()
	}
}
