package mgquerry

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// AppQuery ...
type AppQuery struct {
	Page           int64
	Limit          int64
	SortInterface  interface{}
	Keyword        string
	FromAt         time.Time
	ToAt           time.Time
	Status         string
	SortString     string
	ExpenseTracker ExpenseTracker
}

// CheckLimit ...
func (q *AppQuery) CheckLimit() {
	if q.Limit > 100 || q.Limit < 1 {
		q.Limit = 100
	}
}

// SetDefaultLimit ...
func (q *AppQuery) SetDefaultLimit() {
	if q.Limit == 0 {
		q.Limit = 20
	}
}

// GetFindOptionsWithPage ...
func (q *AppQuery) GetFindOptionsWithPage() *options.FindOptions {
	opts := options.Find()

	// Check limit first
	q.SetDefaultLimit()

	// Add limit & page
	opts.SetLimit(q.Limit).SetSkip(q.Limit * q.Page)

	// Sort
	if q.SortInterface != nil {
		opts.SetSort(q.SortInterface)
	}

	return opts
}
