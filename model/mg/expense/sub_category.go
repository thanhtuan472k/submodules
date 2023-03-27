package expense

import (
	"time"
)

// SubCategory ...
type SubCategory struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	Status       string    `bson:"status"`
	Type         string    `bson:"type"` // match from category
	Category     AppID     `bson:"category"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}

// IsSubCategoryAvailableByTypeIncome ...
func (m SubCategory) IsSubCategoryAvailableByTypeIncome() bool {
	return !m.ID.IsZero() && m.Status == "active" && m.Type == "income"
}
