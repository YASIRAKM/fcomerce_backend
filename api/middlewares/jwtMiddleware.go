package middlewares

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func SetJWTMiddlewares(g *echo.Group) {
	g.Use(
		echojwt.WithConfig(echojwt.Config{
			SigningKey:    []byte("secrete"),
			SigningMethod: "HS512",
			ErrorHandler: func(c echo.Context, err error) error {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired JWT", err)
			},
		}),
	)
}
