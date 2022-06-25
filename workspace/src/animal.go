package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	// variable declaration
	var name, request string
	var cow Animal
	var bird Animal
	var snake Animal

	//data seed
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"

	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"

	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"
	for {
		fmt.Println("Enter the animal name and request")
		r := bufio.NewReader(os.Stdin)
		fmt.Fscanf(r, "%s %s", &name, &request)

		switch name {
		case "cow":
			switch strings.ToLower(request) {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()

			}

		case "bird":
			switch strings.ToLower(request) {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()

			}
		case "snake":
			switch strings.ToLower(request) {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}

		}

	}
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}
