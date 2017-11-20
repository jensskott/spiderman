package parser

import (
	"errors"

	"gopkg.in/yaml.v2"
)

// ParseDefinition to get the service file into a struct
func ParseDefinition(yamlFile []byte, cluster string) (Definition, error) {
	var def Definition

	err := yaml.Unmarshal(yamlFile, &def)
	if err != nil {
		return Definition{}, err
	}

	def.Cluster = cluster

	err = def.validateDefintion()
	if err != nil {
		return Definition{}, err
	}

	return def, nil
}

func (d *Definition) validateDefintion() error {
	if d.Cluster == "" {
		return errors.New("cluster is required")
	} else if d.Service.Name == "" {
		return errors.New("service name is required")
	}

	return nil
}
