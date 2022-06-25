package lead

import (
	"github.com/bosbright/go_fiber_crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var lead []Lead
	db.Find(&lead)
	c.JSON(lead)

}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with this id: %v", id)
	}
	db.Delete(lead)
	c.Send("Lead successfully deleted")

}
