package executor

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

type CommandInput struct {
	Command string `json:"command"`
}

type Executor struct{}

func NewExecutor(e *echo.Echo, executor Executor) {
	e.POST("/run-jq", executor.executeCommand)
}

func (e Executor) executeCommand(c echo.Context) error {
	testJson := "{\"id\": 1234, \"command\": \"hola\"}"

	var input CommandInput
	c.Bind(&input)

	cmd := exec.Command("echo", testJson, "|", "jq", ".")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	fmt.Println(string(output))
	return c.JSON(http.StatusCreated, map[string]string{"result": string(output)})
}
