package cfonts

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed fonts/*.json
var fontsFS embed.FS

type FontSchema struct {
	Name            string              `json:"name"`
	Version         string              `json:"version"`
	Homepage        string              `json:"homepage"`
	Colors          int                 `json:"colors"`
	Lines           int                 `json:"lines"`
	Buffer          []string            `json:"buffer"`
	Letterspace     []string            `json:"letterspace"`
	LetterspaceSize int                 `json:"letterspace_size"`
	Chars           map[string][]string `json:"chars"`
}

// GetFont loads a font from the embedded file system.
func GetFont(name string) (*FontSchema, error) {
	filename := fmt.Sprintf("fonts/%s.json", name)
	data, err := fontsFS.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("font not found: %s", name)
	}

	var schema FontSchema
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("invalid font JSON: %v", err)
	}

	return &schema, nil
}
