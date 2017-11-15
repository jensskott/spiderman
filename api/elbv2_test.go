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
				LoadBalancerArn: aws.String("arn:aws:elasticloadbalancing:us-east-1:444622142361:loadbalancer/app/test-cluster-test-service/b81f5af1ac75b08b"),
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
	assert.Equal(t, "arn:aws:elasticloadbalancing:us-east-1:444622142361:loadbalancer/app/test-cluster-test-service/b81f5af1ac75b08b", testResp)
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
