package iteration

import "strings"

func Repeat(charactor string, times int) string {
	var repeated strings.Builder

	for i := 0; i < times; i++ {
		repeated.WriteString(charactor)
	}
	return repeated.String()
}
