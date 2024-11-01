package utils

import (
	"regexp"
	"strings"
)

// ParseError returns the field that caused the error in a foreign key constraint
// Use exclusively for MySQL errors. It receives the error and returns the field name
func ParseErrorFromForeingKey(err error) (field string) {
	message := err.Error()
	substr := strings.Split(message, "FOREIGN KEY (`")[1]
	field = strings.Split(substr, "`)")[0]

	return
}

// ExtractFKFieldName returns the field that caused the error in a foreign key constraint
// Using regular expressions. It receives the error message and returns the field name
func ExtractFKFieldName(errorMessage string) string {
	re := regexp.MustCompile(`FOREIGN KEY \(` + "`([^`]+)`" + `\)`)
	match := re.FindStringSubmatch(errorMessage)
	if len(match) >= 2 {
		return match[1]
	}
	return ""
}
