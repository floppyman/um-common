package web

import urlpkg "net/url"

func QueryStringBuilder(items map[string]string) string {
	var queryStr = urlpkg.Values{}
	for key, value := range items {
		queryStr.Set(key, value)
	}
	return queryStr.Encode()
}
