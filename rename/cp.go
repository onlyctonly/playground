package main

import (
	"os"
	"strconv"
)

func main() {
	s := 6189546850
	for i := 0; i < 100; i++ {
		os.Link("/Users/jxyu/Downloads/批量上传/pic/base.jpg", "/Users/jxyu/Downloads/批量上传/pic/"+strconv.Itoa(s)+".jpg")
		s++
	}
}
