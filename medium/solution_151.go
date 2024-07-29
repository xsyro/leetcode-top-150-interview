package medium

import (
	"slices"
	"strings"
)

func ReverseWords(s string) string {
	words := slices.DeleteFunc(strings.Split(s, " "), func(s string) bool {
		return s == "" || s == " "
	})
	builder := strings.Builder{}
	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		if i != 0 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}
