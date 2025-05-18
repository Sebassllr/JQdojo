package executor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCommand(t *testing.T) {
	e := echo.New()

	// Sample input
	body := CommandInput{
		Command: ".name",
		Level:   1,
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler
	executor := Executor{}
	if assert.NoError(t, executor.executeCommand(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Alice")
	}

}

func TestExecuteCommand_HTTP(t *testing.T) {
	e := echo.New()
	body := `{"command": ".", "level": 1}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	executor := Executor{}
	err := executor.executeCommand(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Alice")
}
