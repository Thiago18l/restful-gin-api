package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thiago18l/restful-gin-api/src/models"
)

//GET /books
// Get all books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

//POST /books
// Create new books
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

//GET /books/:id
// Find single book
func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

//PATCH /books/:id
//Update a book
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	//validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

//DELETE /book/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
