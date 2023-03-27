package expense

import (
	"time"
)

// IncomeMoney ...
type IncomeMoney struct {
	ID        AppID         `bson:"_id"`
	Category  CategoryShort `bson:"category"` // get list category have type = income
	User      AppID         `bson:"user"`     // get current user login
	Money     float64       `bson:"money"`
	Status    string        `bson:"status"`
	Note      string        `bson:"note"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

// CategoryShort ...
type CategoryShort struct {
	ID   AppID  `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
}

// SubCategoryShort ...
type SubCategoryShort struct {
	ID   AppID  `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
}
