package parser

import (
	"bytes"
	"reflect"
	"testing"
)

// TODO: Add more descriptive and meaningful tests along with comments to describe use case

func TestUnmarshal(t *testing.T) {
	yamlData := []byte(`
        name: John
        description: Example
    `)

	target := T{}

	err := unmarshal(yamlData, &target)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	expected := T{
		Name:        "John",
		Description: "Example",
	}

	if target != expected {
		t.Errorf("Expected %v\nGot: %v", expected, target)
	}
}

func TestStringToByteArr(t *testing.T) {
	data := `
        name: John
        description: Example
	`

	got, _ := stringToByteArr(data)

	expected := []byte(data)

	if !bytes.Equal(expected, got) {
		t.Errorf("Expected %v\nGot: %v", expected, got)
	}
}

func TestParse(t *testing.T) {
	got := T{}
	data := `
        name: John
        description: Example
    `

	err := ParseOnce(data, &got)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expected := T{
		Name:        "John",
		Description: "Example",
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v\nGot: %v", expected, got)
	}
}
