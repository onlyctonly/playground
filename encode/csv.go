package main

import (
	"encoding/csv"
	"os"
	"fmt"
)

func main() {
	f,err:=os.Open("github.com/onlyctonly/goworkspace/playground/encode/test.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	fc:=csv.NewReader(f)
	results,err:=fc.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _,v:=range results[0] {
		fmt.Println(v)
	}
}
