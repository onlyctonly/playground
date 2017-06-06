package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
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
	i:=0
	scanner:=bufio.NewScanner(c)
	for scanner.Scan() {
		line:=scanner.Text()
		if i==0 {
			l:=strings.Fields(line)
			for _,v := range l {
				fmt.Println(v)
			}
			//fmt.Println(l)
		}

		i++
	}
	defer c.Close()
}
