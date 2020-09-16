package test

import (
	"context"
	"demo-user/apptest"
	"demo-user/modules/database"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"demo-user/models"
	"demo-user/util"
)

// Test Create User
type UserCreateTestSuite struct {
	suite.Suite
	e    *echo.Echo
	data models.UserCreatePayload
}

func (suite *UserCreateTestSuite) SetupSuite() {
	// Init server
	suite.e = apptest.InitServer()

	// Clear Data
	removeOldDataUser()

	// Setup payload data
	suite.data = models.UserCreatePayload{
		Name: "hoang",
	}
}

func (suite *UserCreateTestSuite) TearDownSuite() {
	removeOldDataUser()
}

func (suite *UserCreateTestSuite) TestUserCreateSuccess() {
	var (
		payload  = suite.data
		response util.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/users", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	//Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	//Test
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

func (suite *UserCreateTestSuite) TestFail() {
	var (
		payload = models.UserCreatePayload{
			Name: "a",
		}
		response util.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/users", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	//Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	//Test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

// Test List User
type UserListTestSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *UserListTestSuite) SetupSuite() {
	// Init server
	suite.e = apptest.InitServer()

	// Clear Data
	removeOldDataUser()

	// Setup data
	util.HelperUserCreateFake()
}

func (suite *UserListTestSuite) TearDownSuite() {
	removeOldDataUser()
}

func (suite *UserListTestSuite) TestUserListSuccess() {
	var (
		response util.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	//Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	//Test
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// Test Get Transaction BY UserID
type TransactionFindByUserIDTestSuite struct {
	suite.Suite
	e        *echo.Echo
	paramURL string
}

func (suite *TransactionFindByUserIDTestSuite) SetupSuite() {
	// Init server
	suite.e = apptest.InitServer()

	// Clear Data
	removeOldDataUser()

	// Setup param data
	suite.paramURL = util.HelperUserCreateFake()
}

func (suite *TransactionFindByUserIDTestSuite) TearDownSuite() {
	removeOldDataUser()
}

func (suite *TransactionFindByUserIDTestSuite) TestTransactionFindByUserIDSuccess() {
	var (
		response util.Response
	)

	// Setup request
	url := "/users/" + suite.paramURL + "/transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server	
	suite.e.ServeHTTP(rec, req)

	//Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	//Test
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), "Thanh Cong!", response["message"])
}

	func (suite *TransactionFindByUserIDTestSuite) TestTransactionFindByUserIDFailureWithInvalidUserID() {
		var (
			response util.Response
			paramURL = "123"
		)

		// Setup request
		url := "/users/" + paramURL + "/transactions"
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		// Run HTTP server
		suite.e.ServeHTTP(rec, req)

		// Parse
		json.Unmarshal([]byte(rec.Body.String()), &response)

		// Test
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), nil, response["data"])
	}

func (suite *TransactionFindByUserIDTestSuite) TestTransactionFindByUserIDFailureWithNotFoundUser() {
	var (
		response util.Response
		paramURL = "5f24d45125ea51bc11111111"
	)

	// Setup request
	url := "/users" + paramURL + "transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserCreateTestSuite))
	suite.Run(t, new(UserListTestSuite))
	suite.Run(t, new(TransactionFindByUserIDTestSuite))
}

func removeOldDataUser() {
	database.UserCol().DeleteMany(context.Background(), bson.M{})
}
