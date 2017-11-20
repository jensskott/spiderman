package api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jensskott/spiderman/parser"
)

// SearchVpc for the correct VPC for the ALB and SG
func (e *Ec2Implementation) SearchVpc(s *parser.Definition) (string, error) {
	var vpc string

	params := &ec2.DescribeVpcsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(s.Cluster),
				},
			},
		},
	}

	resp, err := e.Svc.DescribeVpcs(params)
	if err != nil {
		return "", err
	}

	for _, v := range resp.Vpcs {
		vpc = *v.VpcId
	}

	return vpc, nil
}

// SearchSubnets for the subnets to add to TargetGroup
func (e *Ec2Implementation) SearchSubnets(vpc string) ([]string, error) {
	var sgs []string
	params := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: aws.StringSlice([]string{vpc}),
			},
		},
	}

	resp, err := e.Svc.DescribeSubnets(params)
	if err != nil {
		return nil, err
	}

	for _, s := range resp.Subnets {
		sgs = append(sgs, *s.SubnetId)
	}

	return sgs, nil
}

// CreateSg for loadbalancer
func (e *Ec2Implementation) CreateSg(s *parser.Definition, vpcID string) ([]string, error) {
	var sg []string

	params := &ec2.CreateSecurityGroupInput{
		Description: aws.String(fmt.Sprintf("Securitygroup for service´s %s loadbalander", s.Service.Name)),
		GroupName:   aws.String(fmt.Sprintf("%s/%s", s.Cluster, s.Service.Name)),
		VpcId:       aws.String(vpcID),
	}

	resp, err := e.Svc.CreateSecurityGroup(params)
	if err != nil {
		return nil, err
	}

	// Add security group id to slice
	sg = append(sg, *resp.GroupId)

	return sg, nil
}
