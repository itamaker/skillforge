package main

import (
	"os"

	"github.com/jonny/current-projects/skillforge/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args[1:]))
}
