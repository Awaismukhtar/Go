package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fName string
		lName string
	}

	var sli = make([]Person, 0, 0)
	var input string
	var p Person
	fmt.Printf("Enter the file Name:\n")
	fmt.Scan(&input)
	f, err := os.Open(input)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	s, e := r.ReadString(10)
	for e == nil {
		fullName := strings.Split(s, " ")
		p.fName = fullName[0]
		p.lName = fullName[1]
		sli = append(sli, p)
		s, e = r.ReadString(10)
	}

	for i := 0; i < len(sli); i++ {
		fmt.Printf("First name: %s, last name: %s\n", sli[i].fName, sli[i].lName)
	}
}
