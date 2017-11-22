package main

import (
	"io/ioutil"
	"log"
	"os"

	"fmt"

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
		resp := createService(def)
		fmt.Println(resp)

	case update.FullCommand():
		def := parseYaml(*file, *cluster)
		resp := updateService(def)
		fmt.Println(resp)
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
