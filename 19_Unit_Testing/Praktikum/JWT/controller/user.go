package controller
import (

	"net/http"
	"strconv"
	"mwr/model"
	"mwr/confiq"
	"mwr/helper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)


func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := confiq.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {


	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user model.User
	if err := confiq.DB.First(&user, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user model.User
	if err := confiq.DB.First(&user, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedUser := new(model.User)
	if err := c.Bind(updatedUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	if err := confiq.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user model.User
	if err := confiq.DB.First(&user, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := confiq.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func LoginController(c echo.Context) error {
    login := model.Login{}
    if err := c.Bind(&login); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    user := model.User{}
    if err := confiq.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    if user.Password != login.Password {
        return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
    }

    key := model.Key{}

	var jwtToken = helper.GenerateJWT(key.Secret, user.ID)

	if jwtToken == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot process data")
	}
		return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "login successful",
        "user":    user.Name,
		"token" : jwtToken,
    })
}

