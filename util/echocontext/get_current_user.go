package echocontext

import (
	"expense-tracker-server/external/mongodb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ...
type User struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// GetCurrentUserID ...
func GetCurrentUserID(c echo.Context) primitive.ObjectID {
	idInterface := c.Get("current_user_id") // TODO: refactor bỏ vào constant
	if idInterface == nil {
		return primitive.NilObjectID
	}
	idUserObject, _ := mongodb.NewIDFromString(idInterface.(string))
	return idUserObject
}

// GetCurrenUserByToken ...
func GetCurrenUserByToken(token interface{}) *User {
	if token == nil {
		return nil
	}

	data, ok := token.(*jwt.Token)
	if !ok {
		return nil
	}
	m, ok := data.Claims.(jwt.MapClaims)
	if !ok || !data.Valid || m == nil {
		return nil
	}
	var user = new(User)

	if m["_id"] != "" {
		user.ID = m["_id"].(string)
	}
	if m["name"] != "" {
		name, ok := m["name"]
		if ok {
			user.Name = name.(string)
		}
	}

	return user
}
