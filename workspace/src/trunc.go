package main

import "fmt"

func main() {
	var y float64
	for i := 0; i < 3; i++ {
		fmt.Print("Enter floating number:")
		fmt.Scan(&y)
		x := int(y)
		fmt.Println(x)
	}
}
