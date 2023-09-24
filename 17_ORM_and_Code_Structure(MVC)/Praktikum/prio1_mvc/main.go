package main
import (
	"prio1_mvc/controller"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)
func main() {
	e := echo.New()
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}