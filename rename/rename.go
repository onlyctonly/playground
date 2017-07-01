package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	dir,err:=os.Open("/Users/jxyu/Downloads/批量上传/pic")
	defer dir.Close()
	if err != nil {
		fmt.Println(err)
	}
	files,_:=dir.Readdir(0)
	s:=6189546850
	for _,file :=range files {
		os.Rename("/Users/jxyu/Downloads/批量上传/pic/" + file.Name(), "/Users/jxyu/Downloads/批量上传/pic/"+strconv.Itoa(s)+".jpg")
		s++
	}
}
