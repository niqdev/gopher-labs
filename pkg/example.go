package example

import (
	"fmt"
	"strings"
)

func Hello(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", fmt.Errorf("empty name")
	}
	return buildMessage(name), nil
}

func buildMessage(name string) string {
	return fmt.Sprintf("hello %s", name)
}
