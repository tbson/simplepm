package vldtutil

import (
	"net/http"
	"net/http/httptest"
	"src/util/localeutil"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Mock payload struct for testing
type MockPayload struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

var e *echo.Echo

func TestMain(m *testing.M) {
	e = echo.New()
	e.Validator = &customValidator{Validator: validator.New()}
	localeutil.Init("en")
	m.Run()
}

// Mock function to initialize the Echo context
func createTestContext(method, path string, body string) echo.Context {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

// Test for ValidatePayload function
func TestValidatePayload(t *testing.T) {
	c := createTestContext(http.MethodPost, "/validate", `{"name": "John", "email": "john@example.com"}`)
	var target MockPayload

	result, err := ValidatePayload(c, target)
	assert.NoError(t, err, "ValidatePayload should not return an error for valid input")
	assert.NotNil(t, result, "Expected non-nil result for valid input")
}

// Test for ValidatePayload with invalid data (missing required fields)
func TestValidatePayload_InvalidData(t *testing.T) {
	c := createTestContext(http.MethodPost, "/validate", `{"name": ""}`)
	var target MockPayload

	_, err := ValidatePayload(c, target)
	assert.Error(t, err, "Expected error for missing required fields")
}

// Test for ValidateUpdatePayload function
func TestValidateUpdatePayload(t *testing.T) {
	c := createTestContext(http.MethodPost, "/validate-update", `{"name": "Alice"}`)
	var target MockPayload

	result, err := ValidateUpdatePayload(c, target)
	assert.NoError(t, err, "ValidateUpdatePayload should not return an error for valid partial data")
	assert.NotNil(t, result, "Expected non-nil result for valid input")
}

// Test for ValidateId function
func TestValidateId(t *testing.T) {
	validId := ValidateId("123")
	assert.Equal(t, uint(123), validId, "Expected valid ID to be parsed correctly")

	invalidId := ValidateId("abc")
	assert.Equal(t, uint(0), invalidId, "Expected invalid ID to return 0")
}

// Test for ValidateIds function
func TestValidateIds(t *testing.T) {
	validIds := ValidateIds("1,2,3")
	assert.Equal(t, []uint{1, 2, 3}, validIds, "Expected list of valid IDs to be parsed correctly")

	invalidIds := ValidateIds("1,abc,3")
	assert.Equal(t, []uint{1, 3}, invalidIds, "Expected invalid IDs to be ignored")
}
