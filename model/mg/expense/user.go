package expense

import (
	"time"
)

// User ...
type User struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	Email        string    `bson:"email,omitempty"`
	Gender       string    `bson:"gender"`
	Phone        string    `bson:"phone"`
	Status       string    `bson:"status"`
	Code         []string  `bson:"code"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
