package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/jnprogrammer/go_fiber/database"
)

type Book struct {
	gorm.Model
	Title  string //'json:"title"'
	Author string // 'json:"author"'
	Rating int    //'json:"rating"'

}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	c.Send("A Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Add Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete Book")
}
