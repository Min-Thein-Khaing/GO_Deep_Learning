package greet

import "strings"

func Hello(name string) string {
	clean := normalLizeName(name)
	return "Hello " + clean
}

func normalLizeName(name string) string {

	n := strings.TrimSpace(name)
	if n == "" {
		return "guest"
	}

	return strings.ToUpper(n)
}
