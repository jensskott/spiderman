package parser

// ServiceDefinition to map from the file
type ServiceDefinition struct {
	Name             string              `yaml:"name"`
	LoadbalancerType string              `yaml:"loadbalancertype"`
	Container        ContainerDefinition `yaml:"container"`
}

type ContainerDefinition struct {
	Protocol string `yaml:"protocol"`
	Port     int64  `yaml:"port"`
}

type Definition struct {
	Service ServiceDefinition `yaml:"service"`
	Cluster string
}
