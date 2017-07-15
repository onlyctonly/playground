package main

import (
	"io/ioutil"
	"strings"
	"os"
	"time"
	"strconv"
)
// CGT_02%mm%_001_CMYK.pdf
// CGT_04%mm%_006_K.pdf
func main() {
	var day string
	var month string
	d:= time.Now().Add(time.Hour*24).Day()
	if d<10 {
		day= "0" + strconv.Itoa(d)
	}else {
		day= strconv.Itoa(d)
	}
	pages,_:=ioutil.ReadDir(".")
	m:=int(time.Now().Add(time.Hour*24).Month())
	if m<10 {
		month="0"+strconv.Itoa(m)
	}else {
		month=strconv.Itoa(m)
	}
	cp:=[]string{"01","04","05","08","09","12","13","16","17","20","21","24","25","28","29","32"}
	bp:=[]string{"02","03","06","07","10","11","14","15","18","19","22","23","26","27","30","31"}
	for _, p:=range pages {
		if strings.Contains(strings.ToLower(p.Name()), "pdf") {
			for _,scp:=range cp {
				if strings.Contains(strings.ToLower(p.Name()), scp) {

					os.Rename(p.Name(), "CGT_" + day + month + "_0"+ scp+ "_CMYK.pdf")
				}
			}
			for _,sbp:=range bp {
				if strings.Contains(strings.ToLower(p.Name()), sbp) {

					os.Rename(p.Name(), "CGT_" + day + month + "_0"+ sbp+ "_K.pdf")
				}
			}

		}
	}

}
