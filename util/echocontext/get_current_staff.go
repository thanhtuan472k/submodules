package echocontext

import (
	"expense-tracker-server/external/mongodb"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Staff ...
type Staff struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// GetCurrentStaffID ...
func GetCurrentStaffID(c echo.Context) (id primitive.ObjectID) {
	token := c.Get("user")
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

// GetCurrenStaffByToken ...
func GetCurrenStaffByToken(token interface{}) *Staff {
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
	var staff = new(Staff)

	if m["_id"] != "" {
		staff.ID = m["_id"].(string)
	}
	if m["name"] != "" {
		name, ok := m["name"]
		if ok {
			staff.Name = name.(string)
		}
	}

	return staff
}
