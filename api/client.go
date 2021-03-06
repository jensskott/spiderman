package api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Ec2Client for connection to EC2
func Ec2Client(region, key, secret string) Ec2Implementation {
	var c Ec2Implementation
	creds := credentials.NewStaticCredentials(key, secret, "")
	c.Session = session.Must(session.NewSession(&aws.Config{Credentials: creds, Region: aws.String(region)}))
	c.Svc = ec2.New(c.Session)
	return c
}

// ElbV2Client for connection to elbv2
func ElbV2Client(region, key, secret string) ElbV2Implementation {
	var c ElbV2Implementation
	creds := credentials.NewStaticCredentials(key, secret, "")
	c.Session = session.Must(session.NewSession(&aws.Config{Credentials: creds, Region: aws.String(region)}))
	c.Svc = elbv2.New(c.Session)
	return c
}

// EcsClient for connection to ecs
func EcsClient(region, key, secret string) EcsImplementation {
	var c EcsImplementation
	creds := credentials.NewStaticCredentials(key, secret, "")
	c.Session = session.Must(session.NewSession(&aws.Config{Credentials: creds, Region: aws.String(region)}))
	c.Svc = ecs.New(c.Session)
	return c
}
