package main

import (
	"fmt"
	"os"
)

func main() {
	//const (
	//	headm="Hello Mr "
	//	headf="Hello Ms "
	//	tail="Best regards\nJia"
	//	mt4=`
	//		Hello {{.Gender}} {{.Name}},
	//		I would like to inform you that your provider account {{}}
	//	`
	//)
	type Email struct {
		Gender string
		Surname string
		data []string
	}
	if os.Args[1] == "mt4" {
		e:=Email{"male", "jxyu", []string{"001","002"}}
		fmt.Print(e.Gender+"\n", e.Surname+"\n", e.data[0], "\n", e.data[1], "\n")
	}
	//template gender surname

}