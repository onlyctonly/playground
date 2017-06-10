package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)
type Person struct {
	Name string
	Phone string
}
func main() {
	session,_:=mgo.Dial("mongodb://localhost:27017/test")
	defer session.Clone()
	session.SetMode(mgo.Monotonic, true)

	c:=session.DB("test").C("people")
	//err:=c.Insert(Person{"xiangyu","6936624325"},Person{"jiajia", "13810375974"})
	result:=Person{}
	err:=c.Find(bson.M{"name":"jiajia"}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
