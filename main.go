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

func getTodos(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var todos []Todo
	db.Find(&todos)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": map[string]interface{}{
			"code":     200,
			"messages": "Success fetching all todo list!",
		},
		"data": todos,
	})
}

func updateTodo(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")

	var todo Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Todo not found",
		})
	}

	if err := c.Bind(&todo); err != nil {
		return err
	}

	db.Save(&todo)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": map[string]interface{}{
			"code":    200,
			"message": "Todo updated successfully!",
		},
	})
}

func deleteTodo(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")

	var todo Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Todo not found",
		})
	}

	db.Delete(&todo)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": map[string]interface{}{
			"code":    200,
			"message": "Todo deleted successfully!",
		},
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
	e.POST("/todos", createTodo)
	e.GET("/todos", getTodos)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
