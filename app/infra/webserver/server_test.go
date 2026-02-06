package webserver_test

import (
	"testing"

	"github.com/joaofilippe/discount-club/app/domain/iservices"
	"github.com/joaofilippe/discount-club/app/infra/webserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	server *webserver.Server
}

type IApplicationMock struct {
	mock.Mock
}

func (a *IApplicationMock) DiscountService() iservices.IDiscount {
	return nil
}

func (a *IApplicationMock) UserService() iservices.IUser {
	return nil
}

func (s *ServerTestSuite) SetupTest() {
	applicationMock := &IApplicationMock{}
	s.server = webserver.New(applicationMock)
}

func (s *ServerTestSuite) TestNewServer() {
	assert.NotNil(s.T(), s.server)
	assert.NotNil(s.T(), s.server.Start(":8080"))
}

func (suite *ServerTestSuite) TestSingletonServer() {
	applicationMock := &IApplicationMock{}

	server1 := webserver.New(applicationMock)
	server2 := webserver.New(applicationMock)

	assert.Equal(suite.T(), server1, server2)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
