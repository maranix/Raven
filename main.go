package main

import (
	"fmt"
	"log"

	m "github.com/maranix/raven/models"
	p "github.com/maranix/raven/parser"
)

func main() {
	mod := m.RootModule{}

	err := p.Parse("test.yaml", &mod)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", mod)
}
