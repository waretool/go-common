package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ErrorSuite struct {
	suite.Suite
	message string
}

func (suite *ErrorSuite) SetupTest() {
	suite.message = "foo bar"
}

func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}

func (suite *ErrorSuite) TestConstant() {
	suite.Equal("one or more of the services is null", ServiceNilError)
	suite.Equal("one or more of the repositories is null", RepositoryNilError)

	suite.Equal("sending http %s request to %v with body %v", HttpRequestInfoMessage)
	suite.Equal("received http reply with status code %d and body %v", HttpResponseInfoMessage)
	suite.Equal("error while sending http %s request to %s due to: %s", HttpErrorMessage)

	suite.Equal("cannot create %s due to: %s", EntityCreateError)
	suite.Equal("cannot get %s due to: %s", EntityGetAllError)
	suite.Equal("cannot get %s with id %d due to: %s", EntityGetError)
	suite.Equal("cannot update %s with id %d due to: %s", EntityUpdateError)
	suite.Equal("no fields of %s with id %d was updated", EntityNoFieldsUpdateWarn)
	suite.Equal("cannot delete %s with id %d due to: %s", EntityDeleteError)
	suite.Equal("cannot find %s with criteria %v due to: %s", EntityFindError)
}

func (suite *ErrorSuite) TestNewRestError() {
	tests := []struct {
		data        *RestError
		expectedErr *RestError
	}{
		{
			NewRestError(500, "foo %s", "bar"),
			&RestError{suite.message, 500},
		},
		{
			NewConflictError("foo %s", "bar"),
			&RestError{suite.message, 409},
		},
		{
			NewBadRequestError("foo %s", "bar"),
			&RestError{suite.message, 400},
		},
		{
			NewUnauthorizedError("foo %s", "bar"),
			&RestError{suite.message, 401},
		},
		{
			NewNotFoundError("foo %s", "bar"),
			&RestError{suite.message, 404},
		},
		{
			NewInternalServerError("foo %s", "bar"),
			&RestError{suite.message, 500},
		},
	}

	for _, test := range tests {
		suite.Equal(test.expectedErr, test.data)
	}
}
