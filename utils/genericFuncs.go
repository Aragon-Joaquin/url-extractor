package utils

import (
	"strings"
)

func ContainsAllValues(input string, substring []string) bool {
	for _, sub := range substring {
		if strings.Contains(input, sub) {
			return true
		}
	}
	return false
}

func CheckTopLevelDomain(input string, domains []string) bool {
	for _, val := range domains {
		level := input[len(input)-len(val):]
		if val == level {
			return true
		}
	}
	return false
}
