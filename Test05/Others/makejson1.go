package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

// import "math"

func main() {
    m := make(map[string]string)
	var input string 
	
	fmt.Printf("Name:")
	fmt.Scanln(&input)
	m ["Name"] = input

	fmt.Printf("Address:")
	fmt.Scanln(&input)
	m ["Address"] = input

	// fmt.Printf("Address:")
	// fmt.Scanln(&p.Address)

	// fmt.Printf("person is: %v", p)

	js, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %v", string(js))
}
