package parser

// ServiceDefinition to map from the file
type ServiceDefinition struct {
	Name             string              `yaml:"name"`
	LoadbalancerType string              `yaml:"loadbalancertype"`
	Count            int64               `yaml:"count"`
	Container        ContainerDefinition `yaml:"container"`
}

type ContainerDefinition struct {
	Protocol          string `yaml:"protocol"`
	Port              int64  `yaml:"port"`
	ContainerPort     int64  `yaml:"containerport"`
	Cpu               int64  `yaml:"cpu"`
	MemoryReservation int64  `yaml:"memoryreservation"`
	Memory            int64  `yaml:"memory"`
	Environment       []Env  `yaml:"environment"`
	Image             string `yaml:"image"`
}

type Definition struct {
	Service ServiceDefinition `yaml:"service"`
	Cluster string
}

type Env struct {
	Name  string
	Value string
}
