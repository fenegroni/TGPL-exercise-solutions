package ch5ex16

import "strings"

func JoinStrings(sep string, elems ...string) string {
	var sb strings.Builder
	for i, e := range elems {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(e)
	}
	return sb.String()
}
