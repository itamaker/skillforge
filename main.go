package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SkillSpec struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Audience    string     `json:"audience"`
	Examples    []string   `json:"examples"`
	Tools       []ToolSpec `json:"tools"`
}

type ToolSpec struct {
	Name        string `json:"name"`
	Command     string `json:"command"`
	Description string `json:"description"`
	Example     string `json:"example"`
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		usage()
		return 2
	}

	switch args[0] {
	case "init":
		return runInit(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand %q\n\n", args[0])
		usage()
		return 2
	}
}

func runInit(args []string) int {
	fs := flag.NewFlagSet("init", flag.ContinueOnError)
	specPath := fs.String("spec", "", "path to a skill spec JSON file")
	outDir := fs.String("out", "", "directory to create")
	force := fs.Bool("force", false, "overwrite an existing output directory")
	fs.SetOutput(os.Stderr)

	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *specPath == "" || *outDir == "" {
		fmt.Fprintln(os.Stderr, "both -spec and -out are required")
		return 2
	}

	spec, err := loadSpec(*specPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := scaffold(spec, *outDir, *force); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("generated skill scaffold in %s\n", *outDir)
	return 0
}

func usage() {
	fmt.Println("skillforge scaffolds OpenClaw-ready skill folders.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  skillforge init -spec examples/skill.json -out /tmp/research-skill")
}

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

func scaffold(spec SkillSpec, outDir string, force bool) error {
	if info, err := os.Stat(outDir); err == nil {
		if !force {
			return fmt.Errorf("output directory %s already exists; use -force to overwrite", outDir)
		}
		if !info.IsDir() {
			return fmt.Errorf("output path %s is not a directory", outDir)
		}
		if err := os.RemoveAll(outDir); err != nil {
			return fmt.Errorf("clear output directory: %w", err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("inspect output directory: %w", err)
	}

	paths := []string{
		outDir,
		filepath.Join(outDir, "bin"),
		filepath.Join(outDir, "examples"),
	}
	for _, path := range paths {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", path, err)
		}
	}

	manifest, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("encode manifest: %w", err)
	}

	files := map[string]string{
		filepath.Join(outDir, "SKILL.md"):             renderSkillMarkdown(spec),
		filepath.Join(outDir, "manifest.json"):        string(manifest) + "\n",
		filepath.Join(outDir, "bin", "README.md"):     renderBinReadme(spec),
		filepath.Join(outDir, "examples", "usage.md"): renderUsageGuide(spec),
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}

	return nil
}

func renderSkillMarkdown(spec SkillSpec) string {
	var builder strings.Builder
	builder.WriteString("# " + spec.Name + "\n\n")
	builder.WriteString(spec.Description + "\n\n")

	if spec.Audience != "" {
		builder.WriteString("## Audience\n\n")
		builder.WriteString(spec.Audience + "\n\n")
	}

	builder.WriteString("## Tools\n\n")
	for _, tool := range spec.Tools {
		builder.WriteString("- `" + tool.Name + "`: " + tool.Description + "\n")
		builder.WriteString("  Command: `" + tool.Command + "`\n")
		if tool.Example != "" {
			builder.WriteString("  Example: `" + tool.Example + "`\n")
		}
	}

	if len(spec.Examples) > 0 {
		builder.WriteString("\n## Example Prompts\n\n")
		for _, example := range spec.Examples {
			builder.WriteString("- " + example + "\n")
		}
	}

	builder.WriteString("\n## Usage\n\n")
	builder.WriteString("Place this directory in your workspace and point your agent platform at `SKILL.md`.\n")
	return builder.String()
}

func renderBinReadme(spec SkillSpec) string {
	var builder strings.Builder
	builder.WriteString("# Tool Commands\n\n")
	builder.WriteString("This folder documents the executable commands expected by the generated skill.\n\n")
	for _, tool := range spec.Tools {
		builder.WriteString("## " + tool.Name + "\n\n")
		builder.WriteString("- Command: `" + tool.Command + "`\n")
		builder.WriteString("- Purpose: " + tool.Description + "\n")
		if tool.Example != "" {
			builder.WriteString("- Example: `" + tool.Example + "`\n")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func renderUsageGuide(spec SkillSpec) string {
	var builder strings.Builder
	builder.WriteString("# Usage Guide\n\n")
	builder.WriteString("Suggested user prompts for the generated skill:\n\n")
	for _, example := range spec.Examples {
		builder.WriteString("- " + example + "\n")
	}
	return builder.String()
}
