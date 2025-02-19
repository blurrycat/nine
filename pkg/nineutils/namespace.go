package nineutils

import (
	"fmt"
	"os"
	"strings"
)

func Namespace() string {
	ns := os.Getenv("NAMESPACE")
	if ns != "" {
		return ns
	}

	display := os.Getenv("DISPLAY")
	if display == "" {
		display = ":0.0"
	}

	parts := strings.Split(display, ".")
	if len(parts) > 1 {
		display = strings.Join(parts[:len(parts)-1], ".")
	}

	display = strings.Replace(display, "/", "_", -1)

	return fmt.Sprintf("/tmp/ns.%s.%s", os.Getenv("USER"), display)
}
