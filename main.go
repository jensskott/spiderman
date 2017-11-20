package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jensskott/spiderman/parser"
)

func main() {
	file, _ := filepath.Abs("./example/service_example.yml")
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	cluster := "test-cluster"

	def, err := parser.ParseDefinition(yamlFile, cluster)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(def)
}
