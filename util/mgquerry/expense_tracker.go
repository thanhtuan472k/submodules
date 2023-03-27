package mgquerry

import (
	"expense-tracker-server/external/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ExpenseTracker ...
type ExpenseTracker struct {
	ID        primitive.ObjectID
	Keyword   string
	Status    string
	Type      string // expense, income
	FromAt    time.Time
	ToAt      time.Time
	CreatedAt time.Time
}

// AssignKeyword ...
func (e ExpenseTracker) AssignKeyword(cond *bson.D) {
	if e.Keyword != "" {
		*cond = append(*cond, bson.E{"searchString", mongodb.GenerateQuerySearchString(e.Keyword)})
	}
}

// AssignStatus ...
func (e ExpenseTracker) AssignStatus(cond *bson.D) {
	if e.Status != "" && e.Status != "all" {
		*cond = append(*cond, bson.E{"status", e.Status})
	}
}

// AssignCategoryType ...
func (e ExpenseTracker) AssignCategoryType(cond *bson.D) {
	if e.Type != "" && e.Type != "all" {
		*cond = append(*cond, bson.E{"type", e.Type})
	}
}

// AssignCategoryID ...
func (e ExpenseTracker) AssignCategoryID(cond *bson.D) {
	if !e.ID.IsZero() {
		*cond = append(*cond, bson.E{"category", e.ID})
	}
}
