package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "Invalid user ID",
		})
	}

	var userFound User
	for _, user := range users {
		if user.Id == userID {
			userFound = user
			break
		}
	}

	if userFound.Id == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success get user by ID",
		"user":    userFound,
	})
}

func DeleteUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "Invalid user ID",
		})
	}

	indexUser := -1
	for i, user := range users {
		if user.Id == userID {
			indexUser = i
			break
		}
	}

	if indexUser == -1 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"error": "User not found",
		})
	}

	users = append(users[:indexUser], users[indexUser+1:]...)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success delete user by ID",
	})
}

func UpdateUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "Invalid user ID",
		})
	}

	var userFound *User
	for i := range users {
		if users[i].Id == userID {
			userFound = &users[i]
			break
		}
	}

	if userFound == nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"error": "User not found",
		})
	}

	updatedUser := new(User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "Invalid request data",
		})
	}

	userFound.Name = updatedUser.Name
	userFound.Email = updatedUser.Email
	userFound.Password = updatedUser.Password

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success update user by ID",
		"user":    userFound,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]any{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	users = []User{
        {Id: 1, Name: "UCUP1", Email: "UU@email.com", Password: "password1"},
        {Id: 2, Name: "UCUP2", Email: "UU@email.com", Password: "password2"},
        {Id: 3, Name: "UCUP3", Email: "UU@email.com", Password: "password3"},
    }
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.POST("/users", CreateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
