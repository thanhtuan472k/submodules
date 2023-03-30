package expense

import (
	"expense-tracker-server/external/constant"
	"time"
)

// Category ...
type Category struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	Status       string    `bson:"status"`
	Type         string    `bson:"type"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}

// IsActiveCategory ...
func (m Category) IsActiveCategory() bool {
	return m.Status == constant.StatusActive
}

// IsCategoryAvailableByTypeIncome ...
func (m Category) IsCategoryAvailableByTypeIncome() bool {
	return !m.ID.IsZero() && m.Status == "active" && m.Type == "income"
}

// IsCategoryAvailableByTypeExpense ...
func (m Category) IsCategoryAvailableByTypeExpense() bool {
	return !m.ID.IsZero() && m.Status == "active" && m.Type == "expense"
}
