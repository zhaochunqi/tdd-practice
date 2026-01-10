package iteration

import "strings"

func Repeat(charactor string, times int) string {
	var repeated strings.Builder

	for range times {
		repeated.WriteString(charactor)
	}
	return repeated.String()
}
