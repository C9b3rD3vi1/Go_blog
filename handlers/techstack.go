package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/utils"
	"github.com/gofiber/fiber/v2"
)

// --- LIST ---
func AdminTechStackList(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	var techStacks []models.TechStack
	if err := database.DB.Order("created_at desc").Find(&techStacks).Error; err != nil {
		return c.Status(500).SendString("Error fetching tech stacks")
	}

	return c.Render("admin/techstacks", fiber.Map{
		"Title":      "Tech Stack Management",
		"Admin":      admin,
		"TechStacks": techStacks,
	})
}

// --- CREATE FORM ---
func AdminNewTechStackForm(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}
	return c.Render("admin/new_techstack", fiber.Map{
		"Title": "Add Tech Stack",
		"Admin": admin,
	})
}

// --- CREATE ACTION ---
func AdminCreateTechStack(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	// Upload image
	iconURL, err := utils.UploadImage(c, "icon")
	if err != nil {
		return c.Status(500).SendString("Error uploading image")
	}

	tech := models.TechStack{
		Name:    c.FormValue("name"),
		IconURL: iconURL,
	}

	if tech.Name == "" {
		return c.Render("admin/new_techstack", fiber.Map{
			"Error": "Tech Stack name is required",
			"Admin": admin,
		})
	}

	if err := database.DB.Create(&tech).Error; err != nil {
		return c.Status(500).SendString("Error saving tech stack")
	}

	return c.Redirect("/admin/techstacks")
}

// --- EDIT FORM ---
func AdminEditTechStackForm(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	var tech models.TechStack
	if err := database.DB.First(&tech, id).Error; err != nil {
		return c.Status(404).SendString("Tech stack not found")
	}

	return c.Render("admin/edit_techstack", fiber.Map{
		"TechStack": tech,
		"Admin":     admin,
	})
}

// --- UPDATE ACTION ---
func AdminUpdateTechStack(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	var tech models.TechStack
	if err := database.DB.First(&tech, id).Error; err != nil {
		return c.Status(404).SendString("Tech stack not found")
	}
	
	iconURL, err := utils.UploadImage(c, "icon")
	if err != nil {
		return c.Status(500).SendString("Error uploading image")
	}

	tech.Name = c.FormValue("name")
	tech.IconURL = iconURL

	if err := database.DB.Save(&tech).Error; err != nil {
		return c.Status(500).SendString("Error updating tech stack")
	}

	return c.Redirect("/admin/techstacks")
}

// --- DELETE ---
func AdminDeleteTechStack(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	if err := database.DB.Delete(&models.TechStack{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting tech stack")
	}

	return c.Redirect("/admin/techstacks")
}
