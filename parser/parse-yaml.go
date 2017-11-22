package parser

import (
	"errors"

	"gopkg.in/yaml.v2"
)

// ParseDefinition to get the service file into a struct
func ParseDefinition(yamlFile []byte, cluster string) (*Definition, error) {
	var def *Definition

	err := yaml.Unmarshal(yamlFile, &def)
	if err != nil {
		return nil, err
	}

	def.Cluster = cluster

	err = def.validateDefinition()
	if err != nil {
		return nil, err
	}

	return def, nil
}

func (d *Definition) validateDefinition() error {
	if d.Cluster == "" {
		return errors.New("cluster is required")
	} else if d.Service.Name == "" {
		return errors.New("service name is required")
	} else if d.Service.Container.MemoryReservation > d.Service.Container.Memory {
		return errors.New("memory needs to be exceeding memory reservation")
	}

	return nil
}
