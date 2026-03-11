package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScaffoldCreatesExpectedFiles(t *testing.T) {
	t.Parallel()

	spec := SkillSpec{
		Name:        "Research Assistant",
		Description: "Summarize papers and search local notes.",
		Audience:    "Applied AI teams.",
		Examples:    []string{"Summarize the latest retrieval notes."},
		Tools: []ToolSpec{
			{
				Name:        "paper-search",
				Command:     "paper-search --query '{{query}}'",
				Description: "Search the local paper index.",
				Example:     "paper-search --query 'rag evaluation'",
			},
		},
	}

	outDir := filepath.Join(t.TempDir(), "research-skill")
	if err := scaffold(spec, outDir, false); err != nil {
		t.Fatalf("scaffold() error = %v", err)
	}

	skillBody, err := os.ReadFile(filepath.Join(outDir, "SKILL.md"))
	if err != nil {
		t.Fatalf("read SKILL.md: %v", err)
	}

	if !strings.Contains(string(skillBody), "paper-search") {
		t.Fatalf("SKILL.md did not include generated tool documentation")
	}

	if _, err := os.Stat(filepath.Join(outDir, "manifest.json")); err != nil {
		t.Fatalf("manifest.json missing: %v", err)
	}
}
