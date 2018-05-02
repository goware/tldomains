# Top-Level Domain Parser

Tiny library to parse the subdomain, domain, and tld extension from a host string.

## Usage

```go
extract, err := tldomains.New("/tmp/tld.cache")
if err != nil {
	fmt.Fprintf(os.Stderr, "error creating cache: %s", err)
}
hostInfo := extract.Parse("mmmm.jello.co.uk")
// hostInfo.Subdomain = "mmmm"
// hostInfo.Root = "jello"
// hostInfo.Suffix = "co.uk"
```


## License

http://mozilla.org/MPL/2.0/
