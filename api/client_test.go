package api

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
)

func TestEc2Client(t *testing.T) {
	c := Ec2Client("us-west-1", "test", "test")
	assert.IsType(t, Ec2Implementation{}, c)
	assert.IsType(t, session.Session{}, *c.Session)
}

func TestElbV2Client(t *testing.T) {
	c := ElbV2Client("us-west-1", "test", "test")
	assert.IsType(t, ElbV2Implementation{}, c)
	assert.IsType(t, session.Session{}, *c.Session)
}

func TestEcsClient(t *testing.T) {
	c := EcsClient("us-west-1", "test", "test")
	assert.IsType(t, EcsImplementation{}, c)
	assert.IsType(t, session.Session{}, *c.Session)
}
