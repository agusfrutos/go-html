package handler

import "github.com/labstack/echo/v4"

type HomeHandler struct{}

func (h HomeHandler) HandleHome(c echo.Context) error {
	return nil
}
