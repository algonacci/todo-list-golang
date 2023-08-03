package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": map[string]interface{}{
			"code":    200,
			"message": "Success fetching the API!",
		},
	})
}

func connectDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(":8080"))
}
