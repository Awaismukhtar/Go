package main

import (
	"fmt"
)

var x int

func printX() {
	fmt.Println(x)
}

func incrementX() {
	x++
}

func main() {
	for x < 60 {
		// We create a race condition by printing and incrementing
		// a shared variable x. A delay in one routine still allows
		// the other to execute, leading to unpredictable output.
		go printX()
		go incrementX()
	}
}
