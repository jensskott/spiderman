package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jensskott/spiderman/parser"
)

func main() {
	file, _ := filepath.Abs("./example/service_example.yml")
	cluster := "test-cluster"

	def, err := parser.ParseDefinition(file, cluster)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(def)
}
