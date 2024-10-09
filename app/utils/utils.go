package utils

import "net/http"

func ExtractPath(response *http.Response) string {
	return response.Request.URL.Path
}

func ValidPath(path string) bool {
	if path == "/" {
		return true
	}
	return false
}
