package api

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/golang/mock/gomock"
	"github.com/jensskott/spiderman/_mocks"
	"github.com/jensskott/spiderman/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateTaskDefinition(t *testing.T) {
	resp := &ecs.RegisterTaskDefinitionOutput{
		TaskDefinition: &ecs.TaskDefinition{
			TaskDefinitionArn: aws.String("arn:aws:ecs:us-west-1:120553098344:task-definition/test-definition:3"),
		},
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mocks.NewMockECSAPI(ctrl)
	mockSvc.EXPECT().RegisterTaskDefinition(gomock.Any()).Return(resp, nil)
	// Create a mock implementation
	e := EcsImplementation{
		Svc: mockSvc,
	}

	def := &parser.Definition{
		Cluster: "test-cluster",
		Service: parser.ServiceDefinition{
			Name:             "test-service",
			LoadbalancerType: "application",
			Container: parser.ContainerDefinition{
				Protocol:          "HTTP",
				Port:              8080,
				ContainerPort:     8080,
				Cpu:               20,
				MemoryReservation: 120,
				Memory:            140,
				Environment: []parser.Env{
					{
						Name:  "test",
						Value: "test-value",
					},
				},
			},
		},
	}

	testResp, err := e.CreateTaskDefinition(def)

	assert.NoError(t, err)

	assert.Equal(t, "arn:aws:ecs:us-west-1:120553098344:task-definition/test-definition:3", testResp)

}

func TestCreateTaskDefinitionError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mocks.NewMockECSAPI(ctrl)
	mockSvc.EXPECT().RegisterTaskDefinition(gomock.Any()).Return(nil, errors.New("I got a error"))
	// Create a mock implementation
	e := EcsImplementation{
		Svc: mockSvc,
	}

	_, err := e.CreateTaskDefinition(&parser.Definition{})

	assert.Error(t, err)
}
