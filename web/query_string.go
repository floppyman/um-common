package web

import urlpkg "net/url"

// QueryStringBuilder converts a map of string key and values into a web url query string
//
//goland:noinspection GoUnusedExportedFunction
func QueryStringBuilder(items map[string]string) string {
	var queryStr = urlpkg.Values{}
	for key, value := range items {
		queryStr.Set(key, value)
	}
	return queryStr.Encode()
}
