package nineutils

import "strings"

func PathForUnixClient(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		return strings.Join(parts[1:], "/")
	}
	return "/"
}
