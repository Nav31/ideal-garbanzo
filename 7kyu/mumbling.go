package codewars

import (
	"strings"
	"bytes"
)

func Accum(s string) string {
	parts := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}
	return strings.Join(parts, "-")
}

func Accum1(s string) string {
	s = strings.ToLower(s)
	var b bytes.Buffer
	for i, r := range s {
		b.WriteString(strings.Title(strings.Repeat(string(r), i+1)))
		if i != len(s)-1 {
			b.WriteString("-")
		}
	}
	return b.String()
}