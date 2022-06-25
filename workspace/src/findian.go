package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Enter the string:\n")
	fmt.Scan(&input)
	var lowerInput = strings.ToLower(input)
	var found = strings.ContainsAny(lowerInput, "ian")
	var startsWith = strings.HasPrefix(lowerInput, "i")
	var endsWith = strings.HasSuffix(lowerInput, "n")
	if found && startsWith && endsWith {
		fmt.Println("Found")
	} else {
		fmt.Printf("Not Found")
	}
}
