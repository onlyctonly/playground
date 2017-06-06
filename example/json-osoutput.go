package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
)

type model struct {
	State    bool
	Pictures []string
}

// What went wrong?
// See the next file for the answer.

func main() {
	m := model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	_, err=os.Stdout.Write(bs)
	if err !=nil {
		log.Fatalln(err)
	}
}