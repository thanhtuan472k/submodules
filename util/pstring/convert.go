package pstring

import "strconv"

// StringToInt64 ...
func StringToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
