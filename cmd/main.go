package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success fetching the API!",
	})
}

func main() {
	e := echo.New()

	e.GET("/", index)

	e.Logger.Fatal(e.Start(":8000"))
}
