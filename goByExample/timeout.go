package main

import (
	"time"
	"fmt"
)

func main() {
	c1:=make(chan string)
	go func() {
		time.Sleep(time.Second*2)
		c1<-`result 1`
	}()


	c2:=make(chan string)
	go func() {
		time.Sleep(time.Second*3)
		c2<-`result 2`
	}()

	select {
	case res:= <-c1:
		fmt.Println(res)
	case <-time.After(time.Second*1):
		fmt.Println("result 1 timed out")
	}

	select {
		case res:= <-c2:
			fmt.Println(res)
		case <-time.After(time.Second*4):
			fmt.Println(`result 2 timed out`)
	}

}