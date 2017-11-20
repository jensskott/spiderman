package api

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	mock "github.com/jensskott/spiderman/_mocks"
	"github.com/jensskott/spiderman/parser"
	"github.com/stretchr/testify/assert"
)

func TestSearchVpc(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeVpcsOutput{
		Vpcs: []*ec2.Vpc{
			{
				VpcId: aws.String("vpc-9654d5ee"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("test-cluster"),
					},
				},
			},
		},
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVpcs(gomock.Any()).Return(resp, nil)

	// Create a mock implementation
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Add the cluster to service defintion for filter testing
	s := &parser.Definition{
		Cluster: "test-cluster",
	}

	// Run a mock of the searchVpc
	testResp, err := e.SearchVpc(s)

	// Make sure it does not return error
	assert.NoError(t, err)

	// Compare respons with what you want to get
	assert.Equal(t, "vpc-9654d5ee", testResp)
}

func TestSearchVpcError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add an error to the controller
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVpcs(gomock.Any()).Return(nil, errors.New("I got a error"))

	// Create a mock implementation
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.SearchVpc(&parser.Definition{})

	// Make sure it returns an error
	assert.Error(t, err)

	// Make sure the response is empty
	assert.Empty(t, testResp)
}

func TestSearchSubnets(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeSubnetsOutput{
		Subnets: []*ec2.Subnet{
			{
				SubnetId: aws.String("subnet-7fd25f34"),
				VpcId:    aws.String("vpc-9654d5ee"),
			},
			{
				SubnetId: aws.String("subnet-b61a1cba"),
				VpcId:    aws.String("vpc-9654d5ee"),
			},
		},
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSubnets(gomock.Any()).Return(resp, nil)

	// Create a mock implementation
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run a mock of the searchSubnets
	testResp, err := e.SearchSubnets("vpc-9654d5ee")
	// Make sure it does not return error
	assert.NoError(t, err)

	// Make sure the slice is two
	assert.Equal(t, 2, len(testResp))

	// Compare respons with what you want to get
	assert.Equal(t, "subnet-7fd25f34", testResp[0])
	assert.Equal(t, "subnet-b61a1cba", testResp[1])
}

func TestSearchSubnetsError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add an error to the controller
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSubnets(gomock.Any()).Return(nil, errors.New("I got a error"))

	// Create a mock implementation
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.SearchSubnets("")

	// Make sure it returns an error
	assert.Error(t, err)

	// Make sure the response is empty
	assert.Empty(t, testResp)
}

func TestCreateSg(t *testing.T) {
	// Create response for mock controller
	resp := &ec2.CreateSecurityGroupOutput{
		GroupId: aws.String("sg-3982694c"),
	}

	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add the response to the controller
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().CreateSecurityGroup(gomock.Any()).Return(resp, nil)

	// Create a mock implementation
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Add the cluster to service defintion for filter testing
	s := &parser.Definition{
		Cluster: "test-cluster",
		Service: parser.ServiceDefinition{
			Name: "test-service",
		},
	}

	// Run a mock of the createSg
	testResp, err := e.CreateSg(s, "vpc-9654d5ee")

	// Make sure it does not return error
	assert.NoError(t, err)

	// Compare respons with what you want to get
	assert.Equal(t, "sg-3982694c", testResp[0])
}

func TestCreateSgError(t *testing.T) {
	// Setup gomock controller with data and close it after run
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock controller and add an error to the controller
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().CreateSecurityGroup(gomock.Any()).Return(nil, errors.New("I got a error"))

	// Create a mock implementation
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.CreateSg(&parser.Definition{}, "")

	// Make sure it returns an error
	assert.Error(t, err)

	// Make sure the response is empty
	assert.Nil(t, testResp)
}
