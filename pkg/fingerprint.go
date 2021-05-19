package sslfingerprint

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Fingerprints struct {
	MD5         string
	SHA1        string
	SHA256      string
	SHA512      string
	Issuer      pkix.Name
	CommonNames []string
}

func GetFingerprint(s string, disableNth bool) (*Fingerprints, error) {

	urlData, _ := url.Parse(s)
	requestedHost := urlData.Host
	certFingerprints := Fingerprints{}

	// This will handle if the client does not have the url and its just the domain name.
	if requestedHost == "" && s == "" {
		requestedHost = "localhost"
	} else if requestedHost == "" {
		requestedHost = s
	}

	req, _ := http.NewRequest("HEAD", "https://"+requestedHost, bytes.NewBufferString(""))
	req.Header.Set("User-Agent", "GoFingerprint/2.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Rewrite an error from the HTTP Client to make it clear of possible causes.
		if strings.Contains(err.Error(), ": no such host") {
			return nil, errors.New("failed to connect or lookup host, check dns and network connection")
		}
		return nil, err
	}

	// To handle the situations where the connection fails.
	if resp != nil {
		tls := resp.TLS
		if tls != nil {
			// Try the first one for simplicity
			cert := tls.PeerCertificates[0]

			certFingerprints.CommonNames = cert.DNSNames
			certFingerprints.Issuer = cert.Issuer

			fingerprintMD5 := strings.ToUpper(fmt.Sprintf("%x", md5.Sum(cert.Raw)))
			if disableNth == false {
				fingerprintMD5 = insertNth(fingerprintMD5, 2)
			}
			certFingerprints.MD5 = fingerprintMD5

			fingerprintSHA1 := strings.ToUpper(fmt.Sprintf("%x", sha1.Sum(cert.Raw)))
			if disableNth == false {
				fingerprintSHA1 = insertNth(fingerprintSHA1, 2)
			}
			certFingerprints.SHA1 = fingerprintSHA1

			fingerprintSHA256 := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256(cert.Raw)))
			if disableNth == false {
				fingerprintSHA256 = insertNth(fingerprintSHA256, 2)
			}
			certFingerprints.SHA256 = fingerprintSHA256

			fingerprintSHA512 := strings.ToUpper(fmt.Sprintf("%x", sha512.Sum512(cert.Raw)))
			if disableNth == false {
				fingerprintSHA512 = insertNth(fingerprintSHA512, 2)
			}
			certFingerprints.SHA512 = fingerprintSHA512
		}
	}

	return &certFingerprints, nil
}

// insertNth inserts a set charter every nth char
// Modified, To use `:` in place of `-`
// @link https://stackoverflow.com/a/33633451/5779200
func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1
	for i, runei := range s {
		buffer.WriteRune(runei)
		if i%n == n1 && i != l1 {
			buffer.WriteRune(':')
		}
	}
	return buffer.String()
}
