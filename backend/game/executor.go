package executor

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os/exec"
	"strings"

	"github.com/Sebassllr/JQdojo/level"
	"github.com/labstack/echo/v4"
)

type CommandInput struct {
	Command string `json:"command"`
	Level   int    `json:"level"`
}

type Executor struct{}

func NewExecutor(e *echo.Echo, executor Executor) {
	e.POST("/run-jq", executor.executeCommand)
}

func (e Executor) execute(command string, levelNumber int) (string, error) {
	levelInput, err := level.GetLevelInput(levelNumber)

	if err != nil {
		return "", err
	}

	var stderr bytes.Buffer
	cmd := exec.Command("jq", command)
	cmd.Stdin = strings.NewReader(levelInput)
	cmd.Stderr = &stderr
	output, err := cmd.Output()

	if err != nil {
		return "", errors.New(stderr.String())
	}

	var jqOutput interface{}
	if err := json.Unmarshal(output, &jqOutput); err != nil {
		return "", err
	}

	result, err := json.Marshal(jqOutput)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (e Executor) executeCommand(c echo.Context) error {
	var input CommandInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	output, err := e.execute(input.Command, input.Level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{"result": string(output)})
}
