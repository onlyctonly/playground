package main

import (
	"fmt"
	"os"
	"syscall"
)

var state syscall.Statfs_t

func main() {
	d, err := os.Open("/Users/jxyu/testLogPath")
	defer d.Close()
	if err != nil {
		fmt.Println(err)
	}

	fs, err := d.Readdir(0)
	for _, v := range fs {
		fmt.Println(v.Name(), " ", v.IsDir())
		if !v.IsDir() {
			err := os.Remove("/Users/jxyu/testLogPath/" + v.Name())
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	syscall.Statfs("/", &state)

	fmt.Println(float32(state.Bfree) * float32(state.Bsize) / 1024.0 / 1024.0 / 1024.0)

}
