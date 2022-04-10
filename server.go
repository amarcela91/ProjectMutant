package main

import (
	"ProjectMutant/internal/connection"
	"ProjectMutant/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.Routes(e)
	connection.ConnectMongo()
	e.Logger.Fatal(e.Start(":1323"))

}
