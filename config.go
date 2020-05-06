package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Load yaml from a file to a struct
func LoadConfig(filename string, doc interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	decoder := yaml.NewDecoder(f)
	return decoder.Decode(doc)
}

// Write config file
func WriteConfig(filename string, doc interface{}) error {
	fmt.Println("Write a sample config file:", filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	data, err := yaml.Marshal(doc)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	return err
}
