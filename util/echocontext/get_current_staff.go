package echocontext

import (
	"expense-tracker-server/external/mongodb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetCurrentStaffID ...
func GetCurrentStaffID(c echo.Context) (id primitive.ObjectID) {
	token := c.Get("staff")
	if token == nil {
		return
	}

	data, ok := token.(*jwt.Token)
	if !ok {
		return
	}

	m, ok := data.Claims.(jwt.MapClaims)
	if ok && data.Valid {
		s, ok := m["_id"].(string)
		if ok && s != "" {
			id = mongodb.ConvertStringToObjectID(s)
		}
	}

	return id
}
