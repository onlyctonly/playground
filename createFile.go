package main

import "os"

func main() {
	f,_:=os.Create("test.txt")
	defer f.Close()
}
