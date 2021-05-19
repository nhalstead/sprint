package main

import (
	"os"

	"github.com/nhalstead/sprint/sprint/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
