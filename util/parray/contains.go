package parray

import "github.com/thoas/go-funk"

// Contains ...
func Contains(in interface{}, ele interface{}) bool {
	return funk.Contains(in, ele)
}

// ContainsStr ...
func ContainsStr(in []string, ele string) bool {
	return funk.ContainsString(in, ele)
}
