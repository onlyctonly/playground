package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Name string
	Gender string
	DoB string
	Age int
}

func main() {
	p:=people{"jane doe", "female", "1990-01-01", 23}
	j,_:=json.Marshal(p)
	fmt.Println("slice of bytes j: ", j) // [123 34 78 97 109 101 34 58 34 106 97 110 101 32 100 111 101 34 44 34 71 101 110 100 101 114 34 58 34 102 101 109 97 108 101 34 44 34 68 111 66 34 58 34 49 57 57 48 45 48 49 45 48 49 34 44 34 65 103 101 34 58 50 51 125]
	fmt.Println("json: ", string(j)) // {"Name":"jane doe","Gender":"female","DoB":"1990-01-01","Age":23}

	//js:= []byte{}

	var data people
	_=json.Unmarshal(j,&data)
	fmt.Println("data: ", data.Name)


	s:=`test string`
	fmt.Println([]byte(s))


}
