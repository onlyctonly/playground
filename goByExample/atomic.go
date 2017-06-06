package main

import (
	"time"
	"fmt"
	"sync/atomic"
)

func main() {
	var n uint64 =0
	for i:=0;i<5;i++{
		go func() {
			for {
				atomic.AddUint64(&n,1)
				time.Sleep(time.Millisecond*2)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(atomic.LoadUint64(&n))
}
