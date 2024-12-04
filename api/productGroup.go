package api

import (
	"fcomerce/api/handler"

	"github.com/labstack/echo/v4"
)

func ProductGroup(g *echo.Group) {

	g.POST("/addColor", handler.AddColor)

}
