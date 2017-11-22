package main

import (
	"fmt"
	"log"

	"github.com/jensskott/spiderman/api"
	"github.com/jensskott/spiderman/parser"
)

func updateService(def *parser.Definition) string {
	ecsClient := api.EcsClient(*region)

	td, err := ecsClient.CreateTaskDefinition(def)
	if err != nil {
		log.Fatal(err)
	}

	service, err := ecsClient.UpdateService(def, td)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Updated service: %s", service)
}
