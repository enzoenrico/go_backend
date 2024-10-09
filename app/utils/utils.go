package utils

import "net/http"

func ExtractPath(req *http.Request) string {
	return req.URL.Path
}

func ValidPath(path string) bool {
	if path == "/" {
		return true
	}
	return false
}
