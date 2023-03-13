package expense

import (
	"expense-tracker-server/internal/config"
	"github.com/golang-jwt/jwt"
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
	Password     string    `json:"password"`
	Status       string    `bson:"status"`
	Code         string    `bson:"code"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}

// GenerateAccessToken ...
func (s User) GenerateAccessToken() string {
	// Init claims
	claims := jwt.MapClaims{
		"_id":   s.ID.Hex(),
		"name":  s.Name,
		"phone": s.Phone,
		"exp":   time.Now().Local().Add(GetTimeExpiredToken()).Unix(), // 6 months
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the actual JWT token
	tokenString, _ := token.SignedString([]byte(config.GetENV().Admin.Auth.SecretKey))

	return tokenString
}
