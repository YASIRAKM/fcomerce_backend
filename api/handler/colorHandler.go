package handler

import (
	"fcomerce/db"
	"fcomerce/model"
	"fcomerce/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AddColor(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userRole := claims["role"]

	// Example: Check if the user is an admin
	if userRole != "admin" {
		return utils.Response(c, http.StatusForbidden, false, "Access denied: only admins can add colors", "")
	}

	color := new(model.ColorModel)
	if err := c.Bind(color); err != nil {
		return utils.Response(c, http.StatusBadRequest, false, "Invalid request payload", "")

	}

	if color.ColorName == "" || color.ColorCode == "" {
		return utils.Response(c, http.StatusBadRequest, false, "Color name and color code is required", "")

	}

	query := "INSERT INTO colors (colorname, colorcode) VALUES (?, ?)"
	result, err := db.DB.Exec(query, color.ColorName, color.ColorCode)
	if err != nil {

		return utils.Response(c, http.StatusOK, false, err.Error(), "")

	}

	return utils.Response(c, http.StatusOK, true, "Color added succesfully", result)

}
