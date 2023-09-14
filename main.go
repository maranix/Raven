package main

import (
	"fmt"
	"log"

	m "github.com/maranix/raven/models"
	p "github.com/maranix/raven/parser"
	"gopkg.in/yaml.v3"
)

func main() {
	mod := m.RootModule{}

	data, err := p.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &mod)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", mod)
}
