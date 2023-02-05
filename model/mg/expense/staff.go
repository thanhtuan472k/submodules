package expense

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// Staff ...
type Staff struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	Email        string    `bson:"email,omitempty"`
	Gender       string    `bson:"gender"`
	Phone        string    `bson:"phone"`
	Password     string    `bson:"password"`
	Status       string    `bson:"status"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}

// GenerateAccessToken ...
func (s Staff) GenerateAccessToken() string {
	// Init claims
	claims := jwt.MapClaims{

		"_id":  s.ID,
		"name": s.Name,
		"exp":  time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, _ := token.SignedString([]byte("AuthSecret"))

	return tokenString
}

// GenerateRefreshToken ...
