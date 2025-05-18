package level

import (
	"encoding/json"
	"testing"
)

func TestGetLevelConfiguration(t *testing.T) {
	lvl, err := GetLevelConfiguration(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if lvl.Title != "Get the name" {
		t.Errorf("expected title to be 'Get the name', got '%s'", lvl.Title)
	}
}

func TestGetLevelInput(t *testing.T) {
	jsonStr, err := GetLevelInput(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var input map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &input)
	if err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if input["name"] != "Alice" {
		t.Errorf("expected 'Alice', got %v", input["name"])
	}
}
