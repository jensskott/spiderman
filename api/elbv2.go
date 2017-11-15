package api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/jensskott/spiderman/parser"
)

func (e *ElbV2Implementation) CreateLb(s *parser.Definition, sg []string, subnets []string) (string, error) {
	var lb string

	params := &elbv2.CreateLoadBalancerInput{
		Name:           aws.String(fmt.Sprintf("%s-%s", s.Cluster, s.Service.Name)),
		SecurityGroups: aws.StringSlice(sg),
		Subnets:        aws.StringSlice(subnets),
		Type:           aws.String(s.Service.LoadbalancerType),
	}

	resp, err := e.Svc.CreateLoadBalancer(params)
	if err != nil {
		return "", err
	}

	for _, l := range resp.LoadBalancers {
		lb = *l.LoadBalancerArn
	}

	return lb, nil
}

func (e *ElbV2Implementation) CreateLbListener(s *parser.Definition, lbArn string, tgArn string) (string, error) {
	var listener string

	params := &elbv2.CreateListenerInput{
		DefaultActions: []*elbv2.Action{
			{
				TargetGroupArn: aws.String(tgArn),
				Type:           aws.String("forward"),
			},
		},
		LoadBalancerArn: aws.String(lbArn),
		Port:            aws.Int64(s.Service.Container.Port),
		Protocol:        aws.String(s.Service.Container.Protocol),
	}

	resp, err := e.Svc.CreateListener(params)
	if err != nil {
		return "", err
	}

	for _, l := range resp.Listeners {
		listener = *l.ListenerArn
	}

	return listener, nil

}

func (e *ElbV2Implementation) CreateLbTargetGroup(s *parser.Definition, vpc string) (string, error) {
	var tg string
	params := &elbv2.CreateTargetGroupInput{
		Name:     aws.String(fmt.Sprintf("%s-%s-tg", s.Cluster, s.Service.Name)),
		Port:     aws.Int64(8080),
		Protocol: aws.String(s.Service.Container.Protocol),
		VpcId:    aws.String(vpc),
	}

	resp, err := e.Svc.CreateTargetGroup(params)
	if err != nil {
		return "", err
	}

	for _, t := range resp.TargetGroups {
		tg = *t.TargetGroupArn
	}

	return tg, nil
}
