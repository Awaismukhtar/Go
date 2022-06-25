package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	var sli = make([]int, 0, 3)
	for {
		fmt.Printf("Enter the number:\n")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		value, error := strconv.Atoi(input)
		if error == nil {
			sli = append(sli, value)
			sort.Slice(sli, func(i, j int) bool {
				return sli[i] < sli[j]
			})
		}
	}
	fmt.Printf("Sorted Arra:\n")
	for _, v := range sli {
		fmt.Println(v)
	}

}
