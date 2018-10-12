package view

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetGame(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"items": "hoge"})
}
