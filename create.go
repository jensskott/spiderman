package main

import (
	"fmt"
	"log"

	"github.com/jensskott/spiderman/api"
	"github.com/jensskott/spiderman/parser"
)

func createService(def *parser.Definition) string {
	ec2Client := api.Ec2Client(*region, *key, *secret)
	vpc, err := ec2Client.SearchVpc(def)
	if err != nil {
		log.Fatal(err)
	}
	subnets, err := ec2Client.SearchSubnets(vpc)
	if err != nil {
		log.Fatal(err)
	}

	sg, err := ec2Client.CreateSg(def, vpc)
	if err != nil {
		log.Fatal(err)
	}

	elbClient := api.ElbV2Client(*region, *key, *secret)

	lb, err := elbClient.CreateLb(def, sg, subnets)
	if err != nil {
		log.Fatal(err)
	}

	tg, err := elbClient.CreateLbTargetGroup(def, vpc)
	if err != nil {
		log.Fatal(err)
	}

	_, err = elbClient.CreateLbListener(def, lb, tg)
	if err != nil {
		log.Fatal(err)
	}

	ecsClient := api.EcsClient(*region, *key, *secret)

	td, err := ecsClient.CreateTaskDefinition(def)
	if err != nil {
		log.Fatal(err)
	}

	service, err := ecsClient.CreateService(def, lb, td)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Created service: %s", service)
}
