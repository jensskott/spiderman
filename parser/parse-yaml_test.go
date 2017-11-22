package parser

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestParseDefinition(t *testing.T) {
	yamlFileStr := `
service:
  name: test-service
  loadbalancertype: network

  container:
    protocol: tcp
    port: 8080
`

	yamlFile := []byte(string(yamlFileStr))

	service, err := ParseDefinition(yamlFile, "test-cluster")
	fmt.Println(service)

	assert.NoError(t, err)

	assert.Equal(t, "test-cluster", service.Cluster)
	assert.Equal(t, "test-service", service.Service.Name)
	assert.Equal(t, "tcp", service.Service.Container.Protocol)
}

func TestParseDefinitionYamlError(t *testing.T) {
	yamlFileStr := "`	container: name"
	yamlFile := []byte(string(yamlFileStr))

	_, err := ParseDefinition(yamlFile, "")

	assert.Error(t, err)

}



func TestParseDefinitionMemoryError(t *testing.T) {
	yamlFileStr := `
service:
  name: test-service
  loadbalancertype: network

  container:
    protocol: tcp
    port: 8080
    memory: 128
    memoryreservation: 256
`
	yamlFile := []byte(string(yamlFileStr))

	_, err := ParseDefinition(yamlFile, "test-cluster")

	assert.Error(t, err)
}
