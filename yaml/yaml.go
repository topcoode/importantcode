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
	// Create an instance of your data structure
	data := MyData{
		Foo: "Hello",
		Bar: 42,
	}

	// Marshal the data structure into YAML
	yamlContent, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Write the YAML content to a file
	err = ioutil.WriteFile("output.yaml", yamlContent, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data written to output.yaml")
}
