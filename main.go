package main

import (
	"fmt"
	"io"
	"log"
	"os"

	m "github.com/maranix/raven/models"
	"gopkg.in/yaml.v3"
)

func main() {
	mod := m.RootModule{}

	file, err := os.Open("test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &mod)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", mod)
}
