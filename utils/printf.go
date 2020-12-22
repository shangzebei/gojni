package utils

import "strings"

func Wp(s string, limit int) string {
	if len(s) < limit {
		return s + strings.Repeat(" ", limit-len(s))
	} else {
		return s[:limit]
	}
}
