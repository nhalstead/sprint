package main

import (
	"fmt"

	"github.com/nhalstead/sprint"
)

func main() {
	test, _ := sprint.GetFingerprint("google.com", false)

	fmt.Println("Sha256 Fingerprint of google.com: ", test.SHA256)
}
