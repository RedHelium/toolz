package toolz

import (
	"strconv"
	"strings"
)

func Concat(args ...string) string {

	var str strings.Builder

	for i := range args {
		str.WriteString(args[i])
	}

	return str.String()
}

// Convert float value (range: 0-255) into a HEX code
func FloatToHex(val float32) string {
	return strconv.FormatInt(int64(val), 16)
}

// Convert integer value (range: 0-255) into a HEX code
func IntToHex(val int64) string {
	return strconv.FormatInt(val, 16)
}
