package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

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

func runDraft(args []string) int {
	fs := flag.NewFlagSet("draft", flag.ContinueOnError)
	briefPath := fs.String("brief", "", "path to a natural-language brief")
	catalogPath := fs.String("catalog", "", "optional path to a JSON tool catalog")
	outPath := fs.String("out", "", "optional output JSON path")
	name := fs.String("name", "", "optional explicit skill name")
	maxTools := fs.Int("max-tools", 3, "maximum number of tools to retrieve from the catalog")
	fs.SetOutput(os.Stderr)

	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *briefPath == "" {
		fmt.Fprintln(os.Stderr, "-brief is required")
		return 2
	}

	brief, err := os.ReadFile(*briefPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("read brief: %w", err))
		return 1
	}

	var catalog []ToolSpec
	if *catalogPath != "" {
		catalog, err = loadToolCatalog(*catalogPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}

	spec := draftSpec(string(brief), catalog, *name, *maxTools)
	body, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("encode drafted spec: %w", err))
		return 1
	}

	if *outPath == "" {
		fmt.Println(string(body))
		return 0
	}

	if err := os.WriteFile(*outPath, append(body, '\n'), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("write drafted spec: %w", err))
		return 1
	}
	fmt.Printf("wrote drafted spec to %s\n", *outPath)
	return 0
}
