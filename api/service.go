package api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/jensskott/spiderman/parser"
)

// CreateService in ECS
func (e *EcsImplementation) CreateService(def *parser.Definition, lb string, taskDefinition string) (string, error) {
	params := &ecs.CreateServiceInput{
		Cluster:      aws.String(def.Cluster),
		DesiredCount: aws.Int64(def.Service.Count),
		LoadBalancers: []*ecs.LoadBalancer{
			{
				LoadBalancerName: aws.String(lb),
			},
		},
		ServiceName:    aws.String(def.Service.Name),
		TaskDefinition: aws.String(taskDefinition),
	}

	resp, err := e.Svc.CreateService(params)
	if err != nil {
		return "", err
	}

	return *resp.Service.ServiceName, nil
}

func (e *EcsImplementation) UpdateService(def *parser.Definition, taskDefinition string) (string, error) {

	params := &ecs.UpdateServiceInput{
		Cluster:        aws.String(def.Cluster),
		DesiredCount:   aws.Int64(def.Service.Count),
		Service:        aws.String(def.Service.Name),
		TaskDefinition: aws.String(taskDefinition),
	}
	resp, err := e.Svc.UpdateService(params)
	if err != nil {
		return "", err
	}

	return *resp.Service.ServiceName, nil
}

// CreateTaskDefinition for the service in ECS
func (e *EcsImplementation) CreateTaskDefinition(def *parser.Definition) (string, error) {
	var env []*ecs.KeyValuePair
	for _, e := range def.Service.Container.Environment {
		kv := &ecs.KeyValuePair{
			Name:  aws.String(e.Name),
			Value: aws.String(e.Value),
		}
		env = append(env, kv)
	}

	params := &ecs.RegisterTaskDefinitionInput{
		Family: aws.String(def.Service.Name),
		ContainerDefinitions: []*ecs.ContainerDefinition{
			{
				Cpu:               aws.Int64(def.Service.Container.Cpu),
				MemoryReservation: aws.Int64(def.Service.Container.MemoryReservation),
				Memory:            aws.Int64(def.Service.Container.Memory),
				Name:              aws.String(fmt.Sprintf("%s-taskdefinition", def.Service.Name)),
				PortMappings: []*ecs.PortMapping{
					{
						ContainerPort: aws.Int64(def.Service.Container.ContainerPort),
					},
				},
				Environment: env,
				Essential:   aws.Bool(true),
				Image:       aws.String(def.Service.Container.Image),
			},
		},
	}

	resp, err := e.Svc.RegisterTaskDefinition(params)
	if err != nil {
		return "", err
	}

	return *resp.TaskDefinition.TaskDefinitionArn, nil
}
