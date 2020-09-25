package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"

	"demo-user/apptest"
	"demo-user/models"
	"demo-user/modules/database"
	"demo-user/utils"
)

// Test UserCreate
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
		payload      = suite.data
		response     utils.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/hoang/Documents/Company/demo-user/schemas/user_create.json")
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/users", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Create JSONLoader from go struct
	documentLoader := gojsonschema.NewGoLoader(response)

	// Validate json response
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	// Test
	assert.Equal(suite.T(), true, result.Valid())
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

func (suite *UserCreateTestSuite) TestUserCreateFailureWithInvalidName() {
	var (
		payload = models.UserCreatePayload{
			Name: "a",
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/users", utils.HelperToIOReader(payload))
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

// Test UserList
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
	utils.HelperUserCreateFake()
}

func (suite *UserListTestSuite) TearDownSuite() {
	removeOldDataUser()
}

func (suite *UserListTestSuite) TestUserListSuccess() {
	var (
		response     utils.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/hoang/Documents/Company/demo-user/schemas/user_detail.json")
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Create JSONLoader from go struct
	documentLoader := gojsonschema.NewGoLoader(response)

	// Validate json response
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	// Test
	assert.Equal(suite.T(), true, result.Valid())
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// Test TransactionFindByUserID
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
	suite.paramURL = utils.HelperUserCreateFake()
}

func (suite *TransactionFindByUserIDTestSuite) TearDownSuite() {
	removeOldDataUser()
}

func (suite *TransactionFindByUserIDTestSuite) TestTransactionFindByUserIDSuccess() {
	var (
		response     utils.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/hoang/Documents/Company/demo-user/schemas/transaction_detail.json")
	)

	// Setup request
	url := "/users/" + suite.paramURL + "/transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Create JSONLoader from go struct
	documentLoader := gojsonschema.NewGoLoader(response)

	// Validate json response
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	// Test
	assert.Equal(suite.T(), true, result.Valid())
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), "Thanh Cong!", response["message"])
}

func (suite *TransactionFindByUserIDTestSuite) TestTransactionFindByUserIDFailureWithInvalidUserID() {
	var (
		response utils.Response
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
		response utils.Response
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

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserCreateTestSuite))
	suite.Run(t, new(UserListTestSuite))
	suite.Run(t, new(TransactionFindByUserIDTestSuite))
}

func removeOldDataUser() {
	database.UserCol().DeleteMany(context.Background(), bson.M{})
}
