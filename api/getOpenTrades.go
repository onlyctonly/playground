package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	c:=&http.Client{
	}
	req,err:=http.NewRequest(http.MethodGet,"http://tradingserver.zulutrade.com/getOpen",nil)
	checkerr(err)
	req.SetBasicAuth("username", "password")
	resp,err:=c.Do(req)
	body,err:=ioutil.ReadAll(resp.Body)
	checkerr(err)
	defer resp.Body.Close()
	fmt.Println(string(body))
	time.Sleep(time.Second*10)
}

func checkerr(e error)  {
	if e != nil {
		fmt.Println(e)
	}
}