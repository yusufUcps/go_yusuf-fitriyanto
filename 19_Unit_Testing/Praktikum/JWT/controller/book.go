package controller

import (
	"net/http"
	"mwr/confiq"
	"mwr/model"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := confiq.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book model.Book
	if err := confiq.DB.First(&book, bookID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by ID",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book model.Book
	if err := confiq.DB.First(&book, bookID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedBook := new(model.Book)
	if err := c.Bind(updatedBook); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book.Judul = updatedBook.Judul
	book.Penulis = updatedBook.Penulis
	book.Penerbit = updatedBook.Penerbit

	if err := confiq.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book model.Book
	if err := confiq.DB.First(&book, bookID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := confiq.DB.Create(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}
