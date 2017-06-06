package main

import "database/sql"
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/tysql")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	var data string
	r, _:=db.Query(`select cust_name from Customers;`)
	defer r.Close()

	for r.Next() {
		_=r.Scan(&data)
		fmt.Println(data)
	}



}
