package api

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/golang/mock/gomock"
	mock "github.com/jensskott/spiderman/_mocks"
	"github.com/jensskott/spiderman/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateLb(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &elbv2.CreateLoadBalancerOutput{
		LoadBalancers: []*elbv2.LoadBalancer{
			{
				LoadBalancerName: aws.String("test-lb"),
			},
		},
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mock.NewMockELBV2API(ctrl)
	mockSvc.EXPECT().CreateLoadBalancer(gomock.Any()).Return(resp, nil)
	// Create a mock implementation
	e := ElbV2Implementation{
		Svc: mockSvc,
	}

	// Add the cluster to service defintion for filter testing
	s := &parser.Definition{
		Service: parser.ServiceDefinition{
			LoadbalancerType: "application",
		},
	}

	sgs := []string{"sg-3982694c"}
	subnets := []string{"subnet-7fd25f34"}

	// Run a mock of the searchVpc
	testResp, err := e.CreateLb(s, sgs, subnets)

	// Make sure it does not return error
	assert.NoError(t, err)

	// Compare respons with what you want to get
	assert.Equal(t, "test-lb", testResp)
}

func TestCreateLbError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add an error to the controller
	mockSvc := mock.NewMockELBV2API(ctrl)
	mockSvc.EXPECT().CreateLoadBalancer(gomock.Any()).Return(nil, errors.New("I got a error"))

	// Create a mock implementation
	e := ElbV2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.CreateLb(&parser.Definition{}, []string{}, []string{})

	// Make sure it returns an error
	assert.Error(t, err)

	// Make sure the response is empty
	assert.Empty(t, testResp)
}

func TestCreateLbListener(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &elbv2.CreateListenerOutput{
		Listeners: []*elbv2.Listener{
			{
				ListenerArn: aws.String("arn:aws:elasticloadbalancing:us-east-1:444622142361:listener/app/test-lb-tf/a443d5379d497457/da735ee060f5d6b1b"),
			},
		},
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mock.NewMockELBV2API(ctrl)
	mockSvc.EXPECT().CreateListener(gomock.Any()).Return(resp, nil)
	// Create a mock implementation
	e := ElbV2Implementation{
		Svc: mockSvc,
	}

	// Add the cluster to service defintion for filter testing
	s := &parser.Definition{
		Service: parser.ServiceDefinition{
			Container: parser.ContainerDefinition{
				Protocol: "HTTP",
				Port:     8080,
			},
		},
	}

	// Run a mock of the searchVpc
	testResp, err := e.CreateLbListener(s, "test-lb-arn", "test-tg-arn")

	// Make sure it does not return error
	assert.NoError(t, err)

	// Compare respons with what you want to get
	assert.Equal(t, "arn:aws:elasticloadbalancing:us-east-1:444622142361:listener/app/test-lb-tf/a443d5379d497457/da735ee060f5d6b1b", testResp)
}

func TestCreateLbListenerError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add an error to the controller
	mockSvc := mock.NewMockELBV2API(ctrl)
	mockSvc.EXPECT().CreateListener(gomock.Any()).Return(nil, errors.New("I got a error"))

	// Create a mock implementation
	e := ElbV2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.CreateLbListener(&parser.Definition{}, "", "")

	// Make sure it returns an error
	assert.Error(t, err)

	// Make sure the response is empty
	assert.Empty(t, testResp)
}

func TestCreateLbTargetGroup(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &elbv2.CreateTargetGroupOutput{
		TargetGroups: []*elbv2.TargetGroup{
			{
				TargetGroupArn: aws.String("arn:aws:elasticloadbalancing:us-east-1:444622142361:targetgroup/test-tg/b3453ad15997f38b"),
			},
		},
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mock.NewMockELBV2API(ctrl)
	mockSvc.EXPECT().CreateTargetGroup(gomock.Any()).Return(resp, nil)
	// Create a mock implementation
	e := ElbV2Implementation{
		Svc: mockSvc,
	}

	// Add the cluster to service defintion for filter testing
	s := &parser.Definition{
		Cluster: "test-cluster",
		Service: parser.ServiceDefinition{
			Name: "test-service",
			Container: parser.ContainerDefinition{
				Protocol: "HTTP",
			},
		},
	}

	// Run a mock of the searchVpc
	testResp, err := e.CreateLbTargetGroup(s, "vpc-9654d5ee")

	// Make sure it does not return error
	assert.NoError(t, err)

	// Compare respons with what you want to get
	assert.Equal(t, "arn:aws:elasticloadbalancing:us-east-1:444622142361:targetgroup/test-tg/b3453ad15997f38b", testResp)
}

func TestCreateLbTargetGroupError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add an error to the controller
	mockSvc := mock.NewMockELBV2API(ctrl)
	mockSvc.EXPECT().CreateTargetGroup(gomock.Any()).Return(nil, errors.New("I got a error"))

	// Create a mock implementation
	e := ElbV2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.CreateLbTargetGroup(&parser.Definition{}, "")

	// Make sure it returns an error
	assert.Error(t, err)

	// Make sure the response is empty
	assert.Empty(t, testResp)
}
