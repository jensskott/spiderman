package parser

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseDefinition(file, cluster string) (Definition, error) {
	var def Definition

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return Definition{}, err
	}

	err = yaml.Unmarshal(yamlFile, &def)
	if err != nil {
		return Definition{}, err
	}

	def.Cluster = cluster

	return def, nil
}
