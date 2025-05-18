package level

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
)

type Level struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Input       interface{} `json:"input"`
	Expected    interface{} `json:"expected"`
}

//go:embed levels/*.json
var levelFiles embed.FS

func GetLevelConfiguration(level int) (Level, error) {
	levelPath := fmt.Sprintf("levels/level%d.json", level)
	file, err := levelFiles.Open(levelPath)
	if err != nil {
		fmt.Println("Error opening embedded file:", err)
		return Level{}, errors.New("no configuration file found")
	}
	defer file.Close()

	byteValue, err := fs.ReadFile(levelFiles, levelPath)
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return Level{}, err
	}

	var levelConfiguration Level
	err = json.Unmarshal(byteValue, &levelConfiguration)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return Level{}, err
	}

	return levelConfiguration, nil
}

func GetLevelInput(level int) (string, error) {
	cfg, err := GetLevelConfiguration(level)
	if err != nil {
		return "", err
	}

	result, err := json.Marshal(cfg.Input)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
