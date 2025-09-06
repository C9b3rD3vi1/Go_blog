package handlers

import (
	"strings"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
)

func AdminListTags(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var tags []models.Tag
    database.DB.Order("name asc").Find(&tags)

    return c.Render("admin/tags", fiber.Map{
        "Tags": tags,
    })
}


func AdminCreateTag(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    name := c.FormValue("name")
    if name == "" {
        return c.Status(400).SendString("Tag name is required")
    }

    tag := models.Tag{Name: strings.TrimSpace(name)}
    if err := database.DB.Create(&tag).Error; err != nil {
        return c.Status(500).SendString("Error creating tag")
    }

    return c.Redirect("/admin/tags")
}


func AdminDeleteTag(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var tag models.Tag
    if err := database.DB.First(&tag, id).Error; err != nil {
        return c.Status(404).SendString("Tag not found")
    }

    if err := database.DB.Delete(&tag).Error; err != nil {
        return c.Status(500).SendString("Error deleting tag")
    }

    return c.Redirect("/admin/tags")
}
