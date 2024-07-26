package easy

import (
	"slices"
	"strings"
)

func lengthOfLastWord(s string) int {
	strArr := strings.Split(s, " ")
	strArr = slices.DeleteFunc(strArr, func(s string) bool {
		return s == " " || s == ""
	})
	return len(strArr[len(strArr)-1])
}
