package main

import (
	"os"

	"github.com/itamaker/skillforge/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args[1:]))
}
