package stringhelper

import (
	"fmt"
	"strings"
)

// Upper is ...
func Upper(text string) string {
	return strings.ToUpper(text)
}

// Concat is ...
func Concat(inputs ...string) string {
	var result string
	for _, input := range inputs {
		result = fmt.Sprintf("%s%s", result, input)
	}
	return result
}
