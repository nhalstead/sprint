package cmd

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/nhalstead/sprint"
	"github.com/spf13/cobra"
)

var (
	useMd5     bool
	useSha1    bool
	useSha256  bool
	useSha512  bool
	disableNth bool
	separator  string

	fingerprintCmd = &cobra.Command{
		Use:   "host [fqdn or ip]",
		Short: "Get ssl fingerprint of a URL",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}

			host := args[0]

			// Default to show SHA 1 if no modes are selected
			if useMd5 == false && useSha1 == false && useSha256 == false && useSha512 == false {
				useSha1 = true
			}

			runeSeparator := getDelimiterFromString(separator)

			// If the URL is HTTP, Do not process it as it does NOT have SSL Cert we are looking for.
			if !strings.HasPrefix(host, "http://") {
				crt, err := sprint.GetFingerprint(host, disableNth)

				if err != nil {
					fmt.Println(err)
					os.Exit(2)
					return
				}

				// Used to insert `,` (separator) between the hashes if multiple are selected.
				alreadyExportedOnce := false

				// MD5 Hash Export
				if useMd5 == true {
					alreadyExportedOnce = true
					fmt.Print(crt.MD5) // MD5 Hash
				}

				// SHA1 Hash Export
				if useSha1 == true {
					if alreadyExportedOnce == true {
						fmt.Printf("%c", runeSeparator)
					}
					alreadyExportedOnce = true
					fmt.Print(crt.SHA1) // SHA1 Hash
				}

				// SHA256 Hash Export
				if useSha256 == true {
					if alreadyExportedOnce == true {
						fmt.Printf("%c", runeSeparator)
					}
					alreadyExportedOnce = true
					fmt.Print(crt.SHA256) // SHA256 Hash
				}

				// SHA512 Hash Export
				if useSha512 == true {
					if alreadyExportedOnce == true {
						fmt.Printf("%c", runeSeparator)
					}
					fmt.Print(crt.SHA512) // SHA512 Hash
				}
				fmt.Println()
			} else {
				os.Exit(2)
			}

		},
	}
)

// getDelimiterFromString Modified from gocsv project
// https://github.com/aotimme/gocsv/blob/ebfb0c7dac7e9bce320a548aa158fd5b1cc9c1c7/cmd/utils.go#L21
func getDelimiterFromString(delimiter string) rune {
	if delimiter == "\\t" {
		return '\t'
	} else if delimiter == "\\n" {
		return '\n'
	}  else if delimiter == "\\r" {
		return '\r'
	} else if len(delimiter) > 0 {
		delimiterRune, _ := utf8.DecodeRuneInString(delimiter)
		return delimiterRune
	}
	return rune(0)
}

func init() {
	fingerprintCmd.Flags().BoolVarP(&useMd5, "md5", "m", false, "return md5 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&useSha1, "sha-1", "1", false, "return sha1 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&useSha256, "sha-256", "2", false, "return sha256 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&useSha512, "sha-512", "5", false, "return sha512 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&disableNth, "disable-nth", "d", false, "disable Nth separator")
	fingerprintCmd.Flags().StringVarP(&separator, "separator", "s", ",", "separator between many hashes")
}
