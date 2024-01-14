package main

import (
	"go-html/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	homeHandler := handler.HomeHandler()

	app.GET("/home", homeHandler.HandleHome)

}
