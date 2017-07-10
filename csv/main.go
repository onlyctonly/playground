package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func checkActivated(s string, ss []string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	// Load all accounts
	f1, err := os.Open("Accounts Report.csv")
	defer f1.Close()
	check(err)

	scanner := bufio.NewScanner(f1)
	f3, err := os.Create("Reports.csv")
	defer f3.Close()
	n := 0
	for scanner.Scan() {

		if n > 0 {
			_, err := io.WriteString(f3, scanner.Text()+"\n")
			check(err)
		}
		n++
	}
	f4, err := os.Open("Reports.csv")
	check(err)

	defer f4.Close()

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f4))
	r.Comma = ';'
	allAccounts, err := r.ReadAll() //how to delete first line, rather than manually
	check(err)
	//load activated accounts
	f2, err := os.Open("activatedAccounts.csv")
	defer f2.Close()
	check(err)
	r2 := csv.NewReader(bufio.NewReader(f2))
	activatedAccounts, err := r2.ReadAll()
	check(err)
	for i := 1; i < len(allAccounts); i++ {
		v, err := strconv.ParseFloat(allAccounts[i][16], 64)
		check(err)
		if v > 0 {
			ac := allAccounts[i][0]
			//start compare with activated accounts
			if checkActivated(ac, activatedAccounts[0]) == false {
				fmt.Println("This account needs to check in CP: ", ac)
			}
		}
	}
	fmt.Println("Check completed.")
	time.Sleep(time.Second * 10)
}
