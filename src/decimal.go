package src

import "strings"

func ReplaceCommaWithDot(value string) string {
	return strings.Replace(value, ",", ".", 1)
}
