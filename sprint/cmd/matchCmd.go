package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/nhalstead/sprint"
	"github.com/spf13/cobra"
)

var (
	matchCmd = &cobra.Command{
		Use:   "match [fqdn or ip] [fingerprint]",
		Short: "Get ssl fingerprint of a URL and check if it matches the given fqdn or ip",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}

			domainURL := args[0]
			fingerprint := args[1]

			crt, err := sprint.GetFingerprint(domainURL, false)

			if err != nil {
				fmt.Println(err)
				os.Exit(2)
				return
			}

			// Create a new slice with set a defined size of 12 (4 fingerprint types, 3 formats)
			certFingerprints := make([]string, 4*3)

			certFingerprints = append(certFingerprints, crt.MD5)
			certFingerprints = append(certFingerprints, crt.SHA1)
			certFingerprints = append(certFingerprints, crt.SHA256)
			certFingerprints = append(certFingerprints, crt.SHA512)

			for _, fingerprint := range certFingerprints {
				certFingerprints = append(certFingerprints, strings.ReplaceAll(fingerprint, ":", " "))
				certFingerprints = append(certFingerprints, strings.ReplaceAll(fingerprint, ":", ""))
			}

			if contains(certFingerprints, fingerprint) {
				os.Exit(0)
			}

			// Does not match given format.
			os.Exit(1)

		},
	}
)

// @link https://stackoverflow.com/a/10485970/5779200
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
