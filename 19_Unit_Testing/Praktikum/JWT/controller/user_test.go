package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mwr/confiq"
	"mwr/helper"
	"mwr/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestCreateUserController_SuccessfulCreation(t *testing.T) {
    e := echo.New()

    userData := `{"name": "John Doe", "email": "johndoe@example.com", "password": "password123"}`
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, CreateUserController(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
    }
}

func TestCreateUserController_ValidationErrors(t *testing.T) {
    // Create a new Echo instance for testing
    e := echo.New()

    // Create a request to the CreateUserController endpoint with invalid user data
    invalidUserData := `{"name": 3, "email": 1, "password": "short"}`
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(invalidUserData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Call the CreateUserController function
    if assert.Error(t, CreateUserController(c)) {
        assert.Equal(t, http.StatusBadRequest, rec.Code)
    }
}

func TestUpdateUserController_Authenticated(t *testing.T) {
    // Create a new Echo instance for testing
    e := echo.New()


    // Create a JWT token for authentication
    key := model.Key{}

    jwtToken := helper.GenerateJWT(key.Secret, 8)

    // Create a request to the UpdateUserController endpoint with JWT authentication
    updatedUserData := `{"name": "Updated Name", "email": "updated@example.com", "password": "newpassword"}`
    req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(updatedUserData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    req.Header.Set(echo.HeaderAuthorization, "Bearer "+jwtToken) // Include the JWT token in the request headers
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Call the UpdateUserController function
    if assert.NoError(t, UpdateUserController(c)) {
        // Check the response status code
        assert.Equal(t, http.StatusOK, rec.Code)

    }
}

func TestUpdateUserController_Unauthenticated(t *testing.T) {
    // Create a new Echo instance for testing
    e := echo.New()

    // Create a request to the UpdateUserController endpoint without JWT authentication
    updatedUserData := `{"name": "Updated Name", "email": "updated@example.com", "password": "newpassword"}`
    req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(updatedUserData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Call the UpdateUserController function
    if assert.NoError(t, UpdateUserController(c)) {
        // Check the response status code for unauthorized access
        assert.Equal(t, http.StatusUnauthorized, rec.Code)

        // You can further check the response body if needed
        // For example, check if it contains an error message
        assert.JSONEq(t, `{"message":"Unauthorized"}`, rec.Body.String())
    }
}

// Similar test functions can be created for other scenarios and error cases.
func TestUpdateUserController(t *testing.T) {
    // Create a new Echo instance for testing
    e := echo.New()

    // Create a mock user and add it to the database
    mockUser := model.User{
        Name:     "Mock User",
        Email:    "mock@example.com",
        Password: "mockpassword",
    }
    if err := confiq.DB.Create(&mockUser).Error; err != nil {
        t.Fatalf("Failed to create mock user: %v", err)
    }
    
    // Create a JWT token for authentication
	key := model.Key{}
    jwtToken := helper.GenerateJWT(key.Secret, mockUser.ID)

    // Create updated user data
    updatedUserData := map[string]interface{}{
        "name":     "Updated Name",
        "email":    "updated@example.com",
        "password": "newpassword",
    }

    // Convert updatedUserData to JSON
    updatedUserDataJSON, err := json.Marshal(updatedUserData)
    if err != nil {
        t.Fatalf("Failed to marshal updatedUserData: %v", err)
    }

    // Create a request to the UpdateUserController endpoint with JWT authentication
    req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/users/%d", mockUser.ID), bytes.NewBuffer(updatedUserDataJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    req.Header.Set(echo.HeaderAuthorization, "Bearer "+jwtToken)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Call the UpdateUserController function
    if assert.NoError(t, UpdateUserController(c)) {
        // Check the response status code
        assert.Equal(t, http.StatusOK, rec.Code)

        // Unmarshal the response JSON to check the message
        var response map[string]interface{}
        if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
            t.Fatalf("Failed to unmarshal response JSON: %v", err)
        }

        // Check the response message
        assert.Equal(t, "success update user", response["message"])

        // Retrieve the updated user from the database
        var updatedUser model.User
        if err := confiq.DB.First(&updatedUser, mockUser.ID).Error; err != nil {
            t.Fatalf("Failed to retrieve updated user from the database: %v", err)
        }

        // Check if user data is updated as expected
        assert.Equal(t, "Updated Name", updatedUser.Name)
        assert.Equal(t, "updated@example.com", updatedUser.Email)
        // You may want to handle password updates securely based on your application's logic
    }
}
