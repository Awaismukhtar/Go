package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	_map := make(map[string]string)
	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter Address:")
	fmt.Scan(&address)

	_map["address"] = address
	_map["name"] = name
	personJson, err := json.Marshal(_map)
	if err != nil {
		fmt.Println(err, "Error occured")
		return
	}
	fmt.Println(string(personJson))
}
