package expense

import (
	"time"
)

// ExpenseMoney ...
type ExpenseMoney struct {
	ID          AppID            `bson:"_id"`
	Category    CategoryShort    `bson:"category"` // get list category have type = expense
	SubCategory SubCategoryShort `bson:"subCategory"`
	User        AppID            `bson:"user"` // get current user login
	Money       float64          `bson:"money"`
	Status      string           `bson:"status"`
	Note        string           `bson:"note"`
	CreatedAt   time.Time        `bson:"createdAt"`
	UpdatedAt   time.Time        `bson:"updatedAt"`
}
