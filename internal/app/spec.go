package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func loadSpec(path string) (SkillSpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SkillSpec{}, fmt.Errorf("read spec: %w", err)
	}

	var spec SkillSpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return SkillSpec{}, fmt.Errorf("decode spec: %w", err)
	}

	if err := validateSpec(spec); err != nil {
		return SkillSpec{}, err
	}
	return spec, nil
}

func loadToolCatalog(path string) ([]ToolSpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read tool catalog: %w", err)
	}

	var direct []ToolSpec
	if err := json.Unmarshal(data, &direct); err == nil && len(direct) > 0 {
		return direct, nil
	}

	var wrapped struct {
		Tools []ToolSpec `json:"tools"`
	}
	if err := json.Unmarshal(data, &wrapped); err != nil {
		return nil, fmt.Errorf("decode tool catalog: %w", err)
	}
	return wrapped.Tools, nil
}

func validateSpec(spec SkillSpec) error {
	if strings.TrimSpace(spec.Name) == "" {
		return errors.New("spec.name is required")
	}
	if strings.TrimSpace(spec.Description) == "" {
		return errors.New("spec.description is required")
	}
	if len(spec.Tools) == 0 {
		return errors.New("spec.tools must contain at least one tool")
	}
	for _, tool := range spec.Tools {
		if strings.TrimSpace(tool.Name) == "" || strings.TrimSpace(tool.Command) == "" {
			return errors.New("every tool requires a name and command")
		}
	}
	return nil
}
