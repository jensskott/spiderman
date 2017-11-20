package api

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

// Ec2Implementation for connecting to EC2
type Ec2Implementation struct {
	Session *session.Session
	Svc     ec2iface.EC2API
}

// ElbV2Implementation for connecting to elbv2
type ElbV2Implementation struct {
	Session *session.Session
	Svc     elbv2iface.ELBV2API
}

type EcsImplementation struct {
	Session *session.Session
	Svc     ecsiface.ECSAPI
}
