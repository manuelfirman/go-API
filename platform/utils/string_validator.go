package utils

import "unicode"

func ContainsNumbers(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}
