package parray

import "github.com/thoas/go-funk"

// UniqueArrayStrings ...
func UniqueArrayStrings(arr []string) []string {
	return funk.UniqString(arr)
}
