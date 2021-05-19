package main

import (
	"os"

	"github.com/nhalstead/ssl-fingerprint/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
