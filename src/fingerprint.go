package main

/**
 * This can Connect to a given domain and Get the SSL Cert's fingerprint.
 * Low Dependicies
 * Good for Automation and making fingerprints on the fly.
 *
 * @author Noah Halstead <nhalstaed00@gmail.com>
 */

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var sha1Export = flag.Bool("sha1", false, "Output SHA1")
var sha256Export = flag.Bool("sha256", false, "Output SHA256")
var sha512Export = flag.Bool("sha512", false, "Output SHA512")
var disableNth = flag.Bool("disableNth", false, "Do not insert `:` into the output hash")
var domainURL = flag.String("domain", "https://localhost", "The Domain to get the SSL Cert Fingerprint from, Default is Localhost")

func main() {
	flag.Parse()
	domainURL := *domainURL

	// Exit if none of the Hashes are Selected.
	if *sha1Export == false && *sha256Export == false && *sha512Export == false {
		os.Exit(1)
	}

	// If the URL is HTTP, Do not process it as it does NOT have SSL Cert we are looking for.
	if !strings.HasPrefix(domainURL, "http://") {
		hashList := getFingerprint(domainURL)

		// Used to insert `,` between the hashes if multible are selected.
		alreadyExportedOnce := false

		// SHA1 Hash Export
		if *sha1Export == true {
			alreadyExportedOnce = true
			fmt.Print(hashList[0]) // SHA1 Hash
		}

		// SHA256 Hash Export
		if *sha256Export == true {
			if alreadyExportedOnce == true {
				fmt.Print(",")
			}
			alreadyExportedOnce = true
			fmt.Print(hashList[1]) // SHA256 Hash
		}

		// SHA512 Hash Export
		if *sha512Export == true {
			if alreadyExportedOnce == true {
				fmt.Print(",")
			}
			fmt.Print(hashList[2]) // SHA512 Hash
		}

	}
}

func getFingerprint(s string) []string {

	urlData, _ := url.Parse(s)
	reqURL := urlData.Host
	hashes := []string{"", "", ""} // Init Blank Values

	// This will handle if the client does not have the url and its just the domain name.
	if reqURL == "" {
		reqURL = s
	}

	req, _ := http.NewRequest("HEAD", "https://"+reqURL, bytes.NewBufferString(""))
	req.Header.Set("User-Agent", "GoFingerprint/1.0.2")

	resp, _ := http.DefaultClient.Do(req)

	// To handle the situations where the connection fails.
	if resp != nil {
		tls := resp.TLS
		if tls != nil {
			// Try the first one for simplicity
			cert := tls.PeerCertificates[0]

			fingerprintSHA1 := strings.ToUpper(string([]byte(fmt.Sprintf("%x", sha1.Sum(cert.Raw)))))
			if *disableNth == false {
				fingerprintSHA1 = insertNth(fingerprintSHA1, 2)
			}
			hashes[0] = fingerprintSHA1

			fingerprintSHA256 := strings.ToUpper(string([]byte(fmt.Sprintf("%x", sha256.Sum256(cert.Raw)))))
			if *disableNth == false {
				fingerprintSHA256 = insertNth(fingerprintSHA256, 2)
			}
			hashes[1] = fingerprintSHA256

			fingerprintSHA512 := strings.ToUpper(string([]byte(fmt.Sprintf("%x", sha512.Sum512(cert.Raw)))))
			if *disableNth == false {
				fingerprintSHA512 = insertNth(fingerprintSHA512, 2)
			}
			hashes[2] = fingerprintSHA512
		}
	}

	return hashes
}

/**
 * Modified, To use `:` in place of `-`
 *
 * @author Oleg <https://stackoverflow.com/users/1209451/oleg>
 * @link https://stackoverflow.com/a/33633451/5779200
 */
func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n1 && i != l1 {
			buffer.WriteRune(':')
		}
	}
	return buffer.String()
}
