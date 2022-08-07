package main

import (
	"fmt"
	"sync"
	"time"
)

var eatWgroup sync.WaitGroup

type fork struct{ sync.Mutex }

type philosopher struct {
	id                  int
	leftFork, rightFork *fork
}

// Goes from thinking to hungry to eating and done eating then starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (_p philosopher) eat() {
	defer eatWgroup.Done()
	for j := 0; j < 3; j++ {
		_p.leftFork.Lock()
		_p.rightFork.Lock()

		display("eating", _p.id)
		time.Sleep(time.Second)

		_p.rightFork.Unlock()
		_p.leftFork.Unlock()

		display("finished eating", _p.id)
		time.Sleep(time.Second)
	}

}

func display(action string, id int) {
	fmt.Printf("Philosopher :%d is %s\n", id+1, action)
}

func main() {
	// How many philosophers and forks

	count := 5

	// Create forks
	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	// Create philospoher, assign them 2 forks and send them to the dining table
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
		eatWgroup.Add(1)
		go philosophers[i].eat()

	}
	eatWgroup.Wait()

}
