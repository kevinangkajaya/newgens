package src

import "strings"

func ReplaceCommaWithDot(value string) string {
	return strings.Replace(value, ",", ".", 1)
}

func GetArrayOfStringSeparatedByEnter(value string) []string {
	var dataToReturn []string
	for _, v := range strings.Split(value, "\n") {
		if v != "" {
			dataToReturn = append(dataToReturn, v)
		}
	}
	return dataToReturn
}

func TrimUnusedCharacters(value string) string {
	return strings.Trim(value, "\n")
}
