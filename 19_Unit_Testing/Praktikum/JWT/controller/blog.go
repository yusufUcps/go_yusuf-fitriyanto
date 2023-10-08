package controller

import (
	"mwr/confiq"
	"mwr/model"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []model.Blog

	if err := confiq.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

func GetBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog model.Blog
	if err := confiq.DB.First(&blog, blogID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := confiq.DB.Preload("User").Find(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by ID",
		"blog":    blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog model.Blog
	if err := confiq.DB.First(&blog, blogID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedBlog := new(model.Blog)
	if err := c.Bind(updatedBlog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	blog.UserID = updatedBlog.UserID
	blog.Judul = updatedBlog.Judul
	blog.Konten = updatedBlog.Konten

	if err := confiq.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Preload("User").Find(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog model.Blog
	if err := confiq.DB.First(&blog, blogID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

func CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Create(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Preload("User").Find(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})

}