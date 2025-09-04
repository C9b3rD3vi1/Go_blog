package handlers

import (
//	"github.com/C9b3rD3vi1/Go_blog/database"
	//"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
)


func ProjectsHandler(c *fiber.Ctx) error {
//	var projects []models.Projects
	//if err := database.DB.Order("created_at desc").Find(&projects).Error; err != nil {
		//return c.Status(500).SendString("Error fetching projects")
		//}

	return c.Render("pages/projects", fiber.Map{
		"Title":    "Projects",
	//	"Projects": projects,
	})
}
