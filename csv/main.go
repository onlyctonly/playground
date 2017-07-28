package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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
	f3, err := os.Create(filepath.Join(os.TempDir(), "Reports.csv"))
	defer f3.Close()
	n := 0
	for scanner.Scan() {

		if n > 0 {
			_, err := io.WriteString(f3, scanner.Text()+"\n")
			check(err)
		}
		n++
	}
	f4, err := os.Open(filepath.Join(os.TempDir(), "Reports.csv"))
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

	// writing results to file start
	f5, err := os.Create("pleaseCheck.txt")
	defer f5.Close()

	// writing results to file end
	for i := 1; i < len(allAccounts); i++ {
		v, err := strconv.ParseFloat(allAccounts[i][16], 64)
		check(err)
		if v > 0 {
			ac := allAccounts[i][0]
			//start compare with activated accounts
			if checkActivated(ac, activatedAccounts[0]) == false {
				_, err := io.WriteString(f5, ac+"\r\n")
				check(err)
			}
		}
	}
	_, err = io.WriteString(f5, "\r\nCheck completed. Please check above broker username (if any) in cp.")
	check(err)
	os.Remove(filepath.Join(os.TempDir(), "Reports.csv"))
}
