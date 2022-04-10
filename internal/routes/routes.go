package routes

import (
	"ProjectMutant/internal/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.POST("/mutant/", controllers.SaveMutant)
	e.GET("/stats", controllers.ObtainMutant)
}
