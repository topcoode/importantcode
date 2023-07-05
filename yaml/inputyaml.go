package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type MyData struct {
	Foo string `yaml:"foo"`
	Bar int    `yaml:"bar"`
}

func main() {
	// Read the YAML file
	yamlContent, err := ioutil.ReadFile("input.yaml")
	if err != nil {
		panic(err)
	}

	// Define a struct variable to hold the unmarshaled data
	var data MyData

	// Unmarshal the YAML content into the struct
	err = yaml.Unmarshal(yamlContent, &data)
	if err != nil {
		panic(err)
	}

	// Access the data in the struct
	fmt.Println("Foo:", data.Foo)
	fmt.Println("Bar:", data.Bar)
}
