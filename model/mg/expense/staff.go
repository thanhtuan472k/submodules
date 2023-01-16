package expense

import "time"

// Staff ...
type Staff struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	Email        string    `bson:"email,omitempty"`
	Gender       string    `bson:"gender"`
	Phone        string    `bson:"phone"`
	Code         string    `bson:"code"`
	Status       string    `bson:"status"`
	IsRoot       bool      `bson:"isRoot"` // IsRoot=true -> Admin
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
