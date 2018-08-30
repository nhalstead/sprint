# SSL Fingerprint

This cli application written in go will connect to the given domain and return the `sha1` fingerprint formatted appropriately.

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
> go run src/fingerprint.go [domain]

### Example Usage
> go run src/fingerprint.go google.com
>
>E3:CF:06:92:EF:33:E6:35:17:34:E5:6E:35:81:52:83:90:06:E1:B7

# Contributors
- nhalstead @ https://github.com/nhalstead
