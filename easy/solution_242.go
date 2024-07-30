package easy

import (
	"slices"
	"strings"
)

func isAnagram(s string, t string) bool {
	if s == t {
		return true
	} else if len(s) != len(t) {
		return false
	}
	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	slices.Sort(sArr)
	slices.Sort(tArr)

	sBuild := strings.Builder{}
	tBuild := strings.Builder{}
	for i := range sArr {
		sBuild.WriteString(sArr[i])
	}
	for i := range tArr {
		tBuild.WriteString(tArr[i])
	}
	return sBuild.String() == tBuild.String()
}
