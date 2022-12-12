package parray

import "github.com/thoas/go-funk"

// Find ...
func Find(arr interface{}, predicate interface{}) interface{} {
	return funk.Find(arr, predicate)
}
