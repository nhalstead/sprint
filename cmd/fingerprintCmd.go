package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/nhalstead/ssl-fingerprint/pkg"
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
			host := args[0]

			// Default to show SHA 1 if no modes are selected
			if useMd5 == false && useSha1 == false && useSha256 == false && useSha512 == false {
				useSha1 = true
			}

			// If the URL is HTTP, Do not process it as it does NOT have SSL Cert we are looking for.
			if !strings.HasPrefix(host, "http://") {
				crt, err := sslfingerprint.GetFingerprint(host, disableNth)

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
						fmt.Printf("%+v", separator)
					}
					alreadyExportedOnce = true
					fmt.Print(crt.SHA1) // SHA1 Hash
				}

				// SHA256 Hash Export
				if useSha256 == true {
					if alreadyExportedOnce == true {
						fmt.Printf("%+v", separator)
					}
					alreadyExportedOnce = true
					fmt.Print(crt.SHA256) // SHA256 Hash
				}

				// SHA512 Hash Export
				if useSha512 == true {
					if alreadyExportedOnce == true {
						fmt.Printf("%+v", separator)
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

func init() {
	fingerprintCmd.Flags().BoolVarP(&useMd5, "md5", "m", false, "return md5 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&useSha1, "sha-1", "1", false, "return sha1 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&useSha256, "sha-256", "2", false, "return sha256 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&useSha512, "sha-512", "5", false, "return sha512 fingerprint")
	fingerprintCmd.Flags().BoolVarP(&disableNth, "disable-nth", "d", false, "disable Nth separator")
	fingerprintCmd.Flags().StringVarP(&separator, "separator", "s", ",", "separator between many hashes")
}
