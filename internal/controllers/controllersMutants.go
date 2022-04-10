package controllers

import (
	"ProjectMutant/internal/models"
	"ProjectMutant/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SaveMutant(c echo.Context) error {
	newMutant := new(models.Mutant)
	if err := c.Bind(newMutant); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	createMutant, err := services.CreateMutant(*newMutant)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return c.JSON(http.StatusOK, createMutant)
}

func ObtainMutant(c echo.Context) error {
	stats, err := services.CalculateStats()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No se logró calcular las estadísticas de verificación de ADN")
	}
	return c.JSON(http.StatusOK, stats)
}
