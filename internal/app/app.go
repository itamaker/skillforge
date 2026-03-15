package app

import (
	"fmt"
	"os"
)

func Run(args []string) int {
	if len(args) == 0 {
		return runTUI()
	}

	switch args[0] {
	case "init":
		return runInit(args[1:])
	case "draft":
		return runDraft(args[1:])
	case "tui", "interactive":
		return runTUI()
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand %q\n\n", args[0])
		usage()
		return 2
	}
}

func usage() {
	fmt.Println("skillforge scaffolds OpenClaw-ready skill folders.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  skillforge                  # launch Bubble Tea TUI")
	fmt.Println("  skillforge init -spec examples/skill.json -out /tmp/research-skill")
	fmt.Println("  skillforge draft -brief examples/brief.md -catalog examples/tools.json -out /tmp/spec.json")
}
