package main

import (
	"fmt"
	"encoding/json"
	"os"
)
//we need initial capital letter to export to json
type Person struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`

}
func main() {
	p1:=Person{"xiangyu", "man", 33}
	fmt.Println("struct: ", p1)
	j, _:=json.Marshal(p1)
	fmt.Println("this is json: ", string(j))

	sliceb:=[]byte(`{"name":"jiajia", "gender":"female", "age": 32}`)
	p2:=Person{}
	json.Unmarshal(sliceb, &p2)
	fmt.Println(p2)

	//write to a file
	f,_:=os.Create("test.txt")
	defer f.Close()
	_,err:=f.WriteString(string(j))
	if err!=nil {
		fmt.Println(err)
	}

}
