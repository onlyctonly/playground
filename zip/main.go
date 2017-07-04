package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"archive/zip"
)

func main() {
	const dir  = "/Users/jxyu/Pictures"
	f,err:=ioutil.ReadDir(dir)
	checkerr(err)

	fzip,_:=os.Create("pic.zip")
	defer fzip.Close()
	w:=zip.NewWriter(fzip)
	defer w.Close()

	for _,file := range f {
		if !file.IsDir() {
			fw,_:=w.Create(file.Name())
			filecontent, err:=ioutil.ReadFile(dir+"/"+file.Name())
			checkerr(err)
			fw.Write(filecontent)
			checkerr(err)
		}
	}

}

func checkerr(e error)  {
	if e!=nil {
		fmt.Println(e)
	}
}