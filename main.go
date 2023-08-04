package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

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

func createTodo(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	var existingTodo Todo
	if err := db.Where("title = ?", todo.Title).First(&existingTodo).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"error": "Todo with the same title already exists",
		})
	}

	result := db.Create(&todo)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to create todo",
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": map[string]interface{}{
			"code":    200,
			"message": "Todo created successfully",
		},
		"data": todo,
	})
}

func main() {
	e := echo.New()

	db := connectDB("root:@tcp(127.0.0.1:3306)/todolist?parseTime=true")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	e.GET("/", index)
	e.POST("/todo", createTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
