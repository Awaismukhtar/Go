package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var numbers []int
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter the integer:")
		fmt.Scan(&input)
		s, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, s)
	}
	BubbleSort(numbers)
	fmt.Println("Sorted Array:", numbers)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		Swap(arr, i)
	}
}

func Swap(arr []int, i int) {
	for j := 0; j < len(arr)-i-1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
