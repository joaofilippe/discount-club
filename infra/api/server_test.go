package api_test

import (
	"testing"

	"github.com/joaofilippe/discount-club/infra/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	server *api.Server
}

func (s *ServerTestSuite) SetupTest() {
	s.server = api.NewServer()
}

func (s *ServerTestSuite) TestNewServer() {
	assert.NotNil(s.T(), s.server)
	assert.NotNil(s.T(), s.server.Start(":8080"))
}

func (suite *ServerTestSuite) TestSingletonServer() {
    server1 := api.NewServer()
    server2 := api.NewServer()

    assert.Equal(suite.T(), server1, server2)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
