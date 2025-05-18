package executor

import (
	"encoding/json"
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
	fmt.Println(input)

	cmdString := fmt.Sprintf(`echo '%s' | jq '%s'`, testJson, input.Command)
	cmd := exec.Command("sh", "-c", cmdString)

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}


	var jqOutput interface{}
	if err := json.Unmarshal(output, &jqOutput); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid jq output"})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"result": jqOutput})
}