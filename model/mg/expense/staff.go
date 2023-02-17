package expense

import (
	"expense-tracker-server/internal/config"
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

// GenerateRefreshToken ...
//func (s Staff) GenerateRefreshToken() string {
//	// Init claims
//	claims := jwt.MapClaims{
//		"_id":  s.ID,
//		"name": s.Name,
//		"exp":  time.Now().Local().Add(GetTimeExpiredToken()).Unix(), // 6 months
//	}
//
//	token := jwt.New(jwt.SigningMethodES256)
//
//	tokenClaim := token.Claims(claims)
//
//	//tokenString, _ := tokenClaim.SignedString([]byte(config.GetENV().Admin.Auth.SecretKey))
//
//	return tokenString
//}

// GetTimeExpiredToken ...
func GetTimeExpiredToken() time.Duration {
	var (
		seconds time.Duration
	)

	seconds = time.Duration(config.GetENV().Admin.Auth.TimeExpiredToken) * time.Second
	return seconds
}
