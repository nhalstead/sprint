# SSL Fingerprint (sprint)

[![Go Report Card](https://goreportcard.com/badge/github.com/nhalstead/ssl-fingerprint)](https://goreportcard.com/report/github.com/nhalstead/ssl-fingerprint)
[![GoDoc](https://godoc.org/github.com/nhalstead/ssl-fingerprint?status.svg)](https://godoc.org/github.com/nhalstead/ssl-fingerprint)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

This cli application written in go will connect to the given domain and return the `md5`/`sha1`/`sha256`/`sha512` fingerprint formatted appropriately.

Simple, Fast, Easy

### Why?

The main purpose of this was for debugging of SSL certs but this package can be used in some automation tasks like cert pinning or validation.

This is a CLI tool and a package.

### How to use the Package

Add the following line to import the package and start using it!

```go
package main

import (
	"fmt"
	// Needed to be imported to work
	"github.com/nhalstead/sprint"
)

func main() {
	test, _ := sprint.GetFingerprint("google.com", false)

	fmt.Println("Sha256 Fingerprint of google.com: ", test.SHA256)
}
```

### How to use the CLI

> sprint host \[domain\] -1

To use SHA256 or SHA512 simply use the `-2` (or `--sha-256`) and `-5` (or `--sha-512`) flags respectively.

> sprint host \[domain\] -2
>
> sprint host \[domain\] -5

#### Example Usage

> sprint host google.com -1
>
> F0\:48\:7A\:59\:65\:34\:33\:F8\:A1\:92\:C6\:C4\:FB\:9A\:CC\:C5\:AD\:0C\:B3\:E2

#### Advanced Usage

If in use for bash scripts you can use `sprint match`. Sprint Match will fetch the fingerprint and return a status using the exit codes.

A Non-zero exit code means that it has failed the check or failed to connect during the process.

This makes it simple to implement a method of cert pinning without the need to use other languages with extra libraries.

#### Help (-h / --help)

```
Usage:
  sprint [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  host        Get ssl fingerprint of a URL
  match       Get ssl fingerprint of a URL and check if it matches the given fqdn or ip

Flags:
  -h, --help   help for sprint
```

```
Usage:
  sprint host [fqdn or ip] [flags]

Flags:
  -d, --disable-nth        disable Nth separator
  -h, --help               help for host
  -m, --md5                return md5 fingerprint
  -s, --separator string   separator between many hashes (default ",")
  -1, --sha-1              return sha1 fingerprint
  -2, --sha-256            return sha256 fingerprint
  -5, --sha-512            return sha512 fingerprint
```

```
Usage:
  sprint match [fqdn or ip] [fingerprint] [flags]

Flags:
  -h, --help   help for match
```
