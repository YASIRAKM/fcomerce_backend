package router

import (
	"fcomerce/api"
	"fcomerce/api/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	productGroup := e.Group("/product")
	productGroup.Use(middleware.CORS())
	middlewares.SetJWTMiddlewares(productGroup)

	api.ProductGroup(productGroup)
	// Define a simple route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	api.MainGroup(e)

	return e

}
