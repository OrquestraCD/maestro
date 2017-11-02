package context

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/rackerlabs/maestro/pkg/config"
)

type Context struct {
	AWSSession *session.Session
	EC2Client  ec2iface.EC2API
	SSMClient  ssmiface.SSMAPI
	Config     config.Config
	done       chan struct{}
}

func New(session *session.Session) *Context {
	return &Context{
		AWSSession: session,
		done:       make(chan struct{}),
		EC2Client:  ec2.New(session),
		SSMClient:  ssm.New(session),
	}
}

func (*Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return c.done
}

func (c *Context) Err() error {
	if _, open := <-c.done; !open {
		return errors.New("context closed")
	}

	return nil
}

func (c *Context) String() string {
	return "context.Context"
}
