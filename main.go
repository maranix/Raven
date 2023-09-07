package main

import (
	"fmt"
	"log"
    "gopkg.in/yaml.v3"
)

type T struct {
	Name        string
	Description string
}

func main() {
	data := `
        name: Test
        description: This is a test for yaml
    `

	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)

	if err != nil {
		log.Fatalln("Couldn't unmarshal data")
	}

	fmt.Printf("%v", t)
}
