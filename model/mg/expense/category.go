package expense

import (
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
