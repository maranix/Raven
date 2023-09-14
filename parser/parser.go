package parser

import (
	"bufio"
	"os"

	"gopkg.in/yaml.v3"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []byte

	for scanner.Scan() {
		line := scanner.Bytes()
		content = append(content, line...)
		content = append(content, '\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func Parse(file string, out interface{}) error {
	data, err := readFile(file)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, out)
	if err != nil {
		return err
	}

	return nil
}
