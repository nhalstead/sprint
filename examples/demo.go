package main

import (
	"fmt"
	// Needed to be imported to work
	"github.com/nhalstead/ssl-fingerprint/pkg"
)

func main() {
	test, _ := sslfingerprint.GetFingerprint("google.com", false)

	fmt.Println("Sha256 Fingerprint of google.com: ", test.SHA256)
}