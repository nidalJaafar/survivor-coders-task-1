package route

import (
	"github.com/labstack/echo"
	"student/controller"
)

func Routes(e *echo.Echo) {
	e.GET("/students", controller.GetAllStudents)
	e.GET("/students/:id", controller.GetStudent)
	e.POST("/students", controller.AddStudent)
	e.DELETE("/students/:id", controller.DeleteStudent)
	e.PUT("/students/:id", controller.UpdateStudent)
	e.PATCH("/students/:id", controller.PatchStudent)
}
