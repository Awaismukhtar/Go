package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type NamedAnimal interface {
	Name() string
}

type AnimalS struct {
	name, food, locomotion, noise string
}
type Cow AnimalS
type Bird AnimalS
type Snake AnimalS

func main() {
	var animals []Animal
	var command, input, animalType string
	for {

		fmt.Printf(">")
		r := bufio.NewReader(os.Stdin)
		fmt.Fscanf(r, "%s %s %s", &command, &input, &animalType)
		command = strings.ToLower(command)
		switch command {
		case "newanimal":
			switch strings.ToLower(animalType) {
			case "cow":
				var _cow = Cow{input, "grass", "walk", "moo"}
				animals = append(animals, _cow)
				fmt.Println("Created it!")
			case "bird":
				var _bird = Bird{input, "worms", "fly", "peep"}
				animals = append(animals, _bird)
				fmt.Println("Created it!")
			case "snake":
				var _snake = Snake{input, "mice", "slither", "hsss"}
				animals = append(animals, _snake)
				fmt.Println("Created it!")
			default:
				fmt.Println("cow,bird and snake are the only options please choose from them")
			}
		case "query":
			switch strings.ToLower(animalType) {
			case "eat":
				_animal := FindAnimal(animals, input, animalType)
				if _animal == nil {
					fmt.Println("Animal not found")
					return
				}
				_animal.Eat()
			case "speak":
				_animal := FindAnimal(animals, input, animalType)
				if _animal == nil {
					fmt.Println("Animal not found")
					return
				}
				_animal.Speak()
			case "move":
				_animal := FindAnimal(animals, input, animalType)
				if _animal == nil {
					fmt.Println("Animal not found")
					return
				}
				_animal.Move()
			default:
				fmt.Println("Please choose from following three e.g eat,move,speak")
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}

func FindAnimal(animals []Animal, input string, animalName string) Animal {
	var animal Animal
	for _, v := range animals {
		namedAnimal := v.(NamedAnimal)
		if namedAnimal.Name() == input {
			animal = v
			break
		}
	}
	return animal
}

func (animal Cow) Eat() {
	fmt.Println(animal.food)
}

func (animal Cow) Speak() {
	fmt.Println(animal.noise)
}

func (animal Cow) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Cow) Name() string {
	return animal.name
}
func (animal Bird) Eat() {
	fmt.Println(animal.food)
}

func (animal Bird) Speak() {
	fmt.Println(animal.noise)
}

func (animal Bird) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Bird) Name() string {
	return animal.name
}
func (animal Snake) Eat() {
	fmt.Println(animal.food)
}

func (animal Snake) Speak() {
	fmt.Println(animal.noise)
}

func (animal Snake) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Snake) Name() string {
	return animal.name
}
