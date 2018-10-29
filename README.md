# SSL Fingerprint

This cli application written in go will connect to the given domain and return the `sha1`/`sha256`/`sha512` fingerprint formatted appropriately.

Simple, Fast, Easy

### Why?
> You can use this package in some automation tasks like cert
>  pinning, Saving the cert pin to a config file.
>  
> Verifying Clients you are connecting to in `cURL`.
>
>  If you have the will power to integrate this,
>  you can do cert pinning by fingerprinting
>  with this by using an if Statement.

### How to use:
> go run src/fingerprint.go -domain [domain] -sha1

To use SHA256 or SHA512 simply use the `-sha256` and `-sha512` flag.
> go run src/fingerprint.go -domain [domain] -sha256

> go run src/fingerprint.go -domain [domain] -sha512

### Example Usage
> go run src/fingerprint.go google.com -sha1
>
> 2B:AE:50:AF:6A:71:43:08:F1:98:A8:23:8A:1E:3A:1A:D2:19:F3:2B

### Advanced Usage and Extras
If you want SHA1 and SHA512 in one request you can combine it into one request
> go run src/fingerprint.go -domain google.com -sha1 -sha512
>
> 2B:AE:50:AF:6A:71:43:08:F1:98:A8:23:8A:1E:3A:1A:D2:19:F3:2B,E8:28: ... :47:4F:5B

To remove the `:` characters from the hash output you can add the flag of `-disableNth` and it will not insert them.
> go run src/fingerprint.go google.com -sha1
>
> 2BAE50AF6A714308F198A8238A1E3A1AD219F32B

**NOTE: If no domain is defined using the Flag, It uses `localhost` as the domain.**

# Contributors
- nhalstead @ https://github.com/nhalstead
