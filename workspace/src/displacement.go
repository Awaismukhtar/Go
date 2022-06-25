package main

import (
	"fmt"
	"time"
)

func main() {
	var acceleration, velocity, displacement float64
	var t float64
	fmt.Printf("Enter Acceleration:\n")
	fmt.Scan(&acceleration)
	fmt.Printf("Enter Initial Velocity:\n")
	fmt.Scan(&velocity)
	fmt.Printf("Enter Displacement:\n")
	fmt.Scan(&displacement)
	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Printf("Enter value for time limit to execute the calculations:\n")
	fmt.Scan(&t)
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("Displacement:", fn(t))
}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	fn :=
		func(t float64) float64 {
			return (0.5 * acceleration * t * t) + (velocity * t) + displacement
		}
	return fn
}
