package midware

import "strings"

func IsTemplate(url string) bool {
	if strings.HasPrefix(url, "/template/") {
		return true
	} else {
		return false
	}
}

func IsAsset(url string) bool {
	if strings.HasPrefix(url, "/asset/") {
		return true
	} else if strings.HasPrefix(url, "/manifest.json") {
		return true
	} else if strings.HasPrefix(url, "/favicon.ico") {
		return true
	} else {
		return false
	}
}
