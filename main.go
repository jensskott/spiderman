package main

import (
	"io/ioutil"
	"log"
	"os"

	"fmt"

	"github.com/jensskott/spiderman/api"
	"github.com/jensskott/spiderman/parser"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("spiderman", "ECS deployment tool")

	create = app.Command("create", "Create Service.")
	update = app.Command("update", "Update service.").Default()

	file    = app.Flag("file", "Your service yaml file.").Required().Short('f').String()
	cluster = app.Flag("cluster", "The cluster you want to deploy to.").Required().Short('c').String()
	region  = app.Flag("region", "Region of the cluster").Default("us-east-1").Short('r').String()
)

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case create.FullCommand():
		def := parseYaml(*file, *cluster)
		fmt.Println(def)

	case update.FullCommand():
		def := parseYaml(*file, *cluster)
		fmt.Println(def)
	}
}

func parseYaml(file, cluster string) *parser.Definition {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	def, err := parser.ParseDefinition(yamlFile, cluster)
	if err != nil {
		log.Fatal(err)
	}

	return def
}

func createService(def *parser.Definition) {
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

	fmt.Println(fmt.Sprintf("Created service: %s", service))
}
