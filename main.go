package main

import (
	"fmt"
	"log"

	p "github.com/maranix/raven/parser"
)

func main() {
	mod := p.T{}
	data := `
        name: Test
        description: This is a test for yaml
    `

	err := p.ParseOnce(data, &mod)

	if err != nil {
		log.Fatalf("Failed to parse data: %v", err)
	}

	fmt.Printf("%v", mod)
}
