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
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	domainURL := os.Args[1]
	fmt.Print(getFingerprint(domainURL))

}

func getFingerprint(s string) string {

	urlData, _ := url.Parse(s)
	reqURL := urlData.Host

	// This will handle if the client does not have the url and its just the domain name.
	if reqURL == "" {
		reqURL = s
	}

	req, _ := http.NewRequest("HEAD", "https://"+reqURL, bytes.NewBufferString(""))
	req.Header.Set("User-Agent", "GoFingerprint/1.0.1")

	resp, _ := http.DefaultClient.Do(req)

	tls := resp.TLS
	if tls != nil {
		// Try the first one for simplicity
		cert := tls.PeerCertificates[0]

		fingerprint := strings.ToUpper(string([]byte(fmt.Sprintf("%x", sha1.Sum(cert.Raw)))))
		fingerprint = insertNth(fingerprint, 2)
		return fingerprint
		// Do something with the signature
	}
	return ""
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
