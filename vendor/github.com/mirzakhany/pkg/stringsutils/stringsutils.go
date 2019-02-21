package stringsutils

import (
	"encoding/json"
	"strings"
)

// ReplaceMsg replace a set of placeholders with values
func ReplaceMsg(message string, placeHolders []string, values []string) string {
	for i, k := range placeHolders {
		message = strings.Replace(message, k, values[i], -1)
	}
	return message
}

// StringifyJson stringify a json
func StringifyJson(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// MaxLen return max len of a string array
func MaxLen(in []string) int {
	maxLen := 1
	for _, i := range in {
		li := len(i)
		if li > maxLen {
			maxLen = li
		}
	}
	return maxLen
}
