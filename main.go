package main

import (
	"fmt"

	"github.com/bosbright/go_fiber_crm/database"
	"github.com/bosbright/go_fiber_crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/vi/lead/", lead.GetLeads)
	app.Get("/api/vi/lead/:id", lead.GetLead)
	app.Post("/api/vi/lead/", lead.NewLead)
	app.Delete("/api/vi/lead/:id", lead.DeleteLead)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()

}
