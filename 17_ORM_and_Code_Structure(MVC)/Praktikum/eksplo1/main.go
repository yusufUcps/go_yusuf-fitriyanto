package main
import (
	"ekplo1/controller"
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

	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)


	e.Logger.Fatal(e.Start(":8000"))
}