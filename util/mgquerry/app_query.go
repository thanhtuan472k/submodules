package mgquerry

import "time"

// AppQuery ...
type AppQuery struct {
	Page           int64
	Limit          int64
	SortInterface  interface{}
	Keyword        string
	FromAt         time.Time
	ToAt           time.Time
	Status         string
	ExpenseTracker ExpenseTracker
}
