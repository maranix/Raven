package parser

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// TODO: Move this to root and centralize type definitions or maybe create an other package?
// An example model type
type T struct {
	Name        string
	Description string
}

// Parses a []byte into an interface
//
// return an error in case of failure
func unmarshal(data []byte, in interface{}) error {
	return yaml.Unmarshal(data, &in)
}

// Converts a string into a []byte after validating and removing non UTF-8 characters.
//
// Since this is an internal function there is no need for returning an error here other than maintaining code consistency.
func stringToByteArr(data string) ([]byte, error) {
	sanitize := strings.ToValidUTF8(data, "")

	return []byte(sanitize), nil
}

// Unmarshals and parses given string into interface only once
func ParseOnce(data string, in interface{}) error {
	byteData, _ := stringToByteArr(data)

	err := unmarshal(byteData, &in)

	if err != nil {
		return err
	}

	return nil
}
