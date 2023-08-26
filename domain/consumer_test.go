package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConsumerSuite struct {
	suite.Suite
}

func TestConsumerSuite(t *testing.T) {
	suite.Run(t, new(ConsumerSuite))
}

func (suite *ConsumerSuite) TestConsumerInterface() {
	suite.Implementsf((*Consumer)(nil), new(fakeConsumer), "Consumer interface is changed")
}
