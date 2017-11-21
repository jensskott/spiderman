package main

import (
	"io/ioutil"
	"log"

	"fmt"

	"github.com/jensskott/spiderman/api"
	"github.com/jensskott/spiderman/parser"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	create = kingpin.Command("create", "Create Service.")
	update = kingpin.Command("update", "Update service.").Default()

	file    = kingpin.Arg("file", "Your service yaml file.").Required().String()
	cluster = kingpin.Arg("cluster", "The cluster you want to deploy to.").Required().String()
	region  = kingpin.Arg("region", "Region of the cluster").Default("us-east-1").Required().String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	yamlFile, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	def, err := parser.ParseDefinition(yamlFile, *cluster)
	if err != nil {
		log.Fatal(err)
	}

	ec2Client := api.Ec2Client(*region)
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

	elbClient := api.ElbV2Client(*region)

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

	ecsClient := api.EcsClient(*region)

	td, err := ecsClient.CreateTaskDefinition(def)
	if err != nil {
		log.Fatal(err)
	}

	service, err := ecsClient.CreateService(def, lb, td)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Created service: %s"), service)
}
