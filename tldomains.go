//go:generate becky -lib=false -wrap=bundle -var=_ tldomains.dat
package tldomains

import "strings"

var DomainsMap = make(map[string]struct{}, 0)

type asset struct {
	Name    string
	Content string
	etag    string
}

func bundle(a asset) asset {
	list := strings.Split(a.Content, "\n")
	for _, item := range list {
		if item == "" || strings.HasPrefix(item, "//") {
			continue
		}
		DomainsMap[item] = struct{}{}
	}
	return a
}

func ParseHost(host string) (string, string, error) {
	return "", "", nil
}
