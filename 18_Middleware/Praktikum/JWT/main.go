package main
import (
	"mwr/controller"
	"mwr/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware" 
)
func main() {
	key := model.Key{}
	e := echo.New()
	e.Use(middleware.Logger())
	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(key.Secret)))
	auth.GET("/users", controller.GetUsersController)
	auth.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.POST("/users/login", controller.LoginController)
	auth.DELETE("/users/:id", controller.DeleteUserController)
	auth.PUT("/users/:id", controller.UpdateUserController)

	auth.GET("/books", controller.GetBooksController)
	auth.GET("/books/:id", controller.GetBookController)
	auth.POST("/books", controller.CreateBookController)
	auth.DELETE("/books/:id", controller.DeleteBookController)
	auth.PUT("/books/:id", controller.UpdateBookController)

	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)


	e.Logger.Fatal(e.Start(":8000"))
}