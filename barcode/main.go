package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const head string = "0;"
	const tail string = ";2;2;4;0;000000;1\n"

	f, err := os.Open("tk.txt")
	checkerr(err)

	f2, _ := os.Create("import.csv")
	checkerr(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(f2, head+scanner.Text()+tail)
	}
	defer f.Close()

}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
