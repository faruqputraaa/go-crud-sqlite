package main

import "github.com/labstack/echo/v4"

//endpoint
func Route(e *echo.Echo) {
	user := e.Group("students")
	user.GET("", GetStudents)
	user.GET("/:id", GetStudent)
	user.POST("", CreateStudent)
	user.PUT("/:id", UpdateStudent)
	user.DELETE("/:id", DeleteStudent)
}
