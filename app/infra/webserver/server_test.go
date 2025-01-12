package webserver_test

import (
	"testing"

	"github.com/joaofilippe/discount-club/app/infra/webserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	server *webserver.Server
}

func (s *ServerTestSuite) SetupTest() {
	s.server = webserver.New()
}

func (s *ServerTestSuite) TestNewServer() {
	assert.NotNil(s.T(), s.server)
	assert.NotNil(s.T(), s.server.Start(":8080"))
}

func (suite *ServerTestSuite) TestSingletonServer() {
	server1 := webserver.New()
	server2 := webserver.New()

	assert.Equal(suite.T(), server1, server2)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
