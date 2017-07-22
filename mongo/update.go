package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, _ := mgo.Dial("mongodb://localhost:27017/test")
	defer session.Clone()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	//err:=c.Insert(Person{"xiangyu","6936624325"},Person{"jiajia", "13810375974"})
	err := c.Update(bson.M{"name": "jiajia"}, bson.M{"$set": bson.M{"phone": "138"}})
	if err != nil {
		fmt.Println(err)
	}
}
