package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	x = x + 1
}

func main() {
	//global var initialised as one
	x = 1
	var wg sync.WaitGroup
	wg.Add(1)
	go increment(&wg)
	wg.Add(1)
	go increment(&wg)
	wg.Wait()
	fmt.Println(x)
}
