package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
)

// AdminDashboard renders the admin dashboard
// AdminDashboard renders the admin dashboard
func AdminDashboard(c *fiber.Ctx) error {
    // Get the current logged-in user
    admin := config.GetCurrentUser(c)
    if admin == nil {
        return c.Redirect("/admin/login")
    }

    // Ensure user is actually an admin
    if !admin.IsAdmin {
        return c.SendStatus(fiber.StatusForbidden)
    }

    // fetch projects, posts, services, and users
    var posts []models.Post
    var projects []models.Projects   // ⚡ not Projects (should match your struct name)
    var services []models.Services   // ⚡ same here
    var users []models.User

    database.DB.Order("created_at desc").Find(&posts)
    database.DB.Order("created_at desc").Find(&projects)
    database.DB.Order("created_at desc").Find(&users)
    database.DB.Order("created_at desc").Find(&services)

    return c.Render("admin/dashboard", fiber.Map{
        "Title":    "Admin Dashboard",
        "Admin":    admin.Username, // pass name, or whole struct if needed
        "Posts":    posts,
        "Projects": projects,
        "Services": services,
        "Users":    users,
    })
}


// AdminPostList renders the admin post list
func AdminPostList(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		return c.Status(500).SendString("Error fetching posts")
	}

	return c.Render("admin/posts", fiber.Map{
		"Title": "Admin Post List",
		"Admin": admin,
		"Posts": posts,
	})
}

// AdminEditPostForm renders the admin post edit form
func AdminEditPostForm(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	var blogpost models.BlogPost
	if err := database.DB.First(&blogpost, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
	}

	return c.Render("admin/edit", fiber.Map{
		"Post": blogpost,
	})
}


// AdminUpdatePost updates a post
func AdminUpdatePost(c *fiber.Ctx) error {
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Redirect("/admin/login")
	}
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	var blogpost models.BlogPost
	if err := database.DB.First(&blogpost, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
	}

	// update fields
	blogpost.Title = c.FormValue("title")
	blogpost.Content = c.FormValue("content")
	blogpost.Slug = c.FormValue("slug")
	blogpost.ImageURL = c.FormValue("image")
	blogpost.Tags = c.FormValue("tags")
	blogpost.Author = admin.(models.User).Username

	if blogpost.Title == "" || blogpost.Slug == "" {
		return c.Render("admin/edit", fiber.Map{
			"Post":  blogpost,
			"Error": "Title and Slug are required",
		})
	}

	if err := database.DB.Save(&blogpost).Error; err != nil {
		return c.Status(500).SendString("Error updating post")
	}

	return c.Redirect("/admin/posts")
}

// AdminDeletePost handles post deletion
func AdminDeletePost(c *fiber.Ctx) error {
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Redirect("/admin/login")
	}
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	var blogpost models.BlogPost
	if err := database.DB.First(&blogpost, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
	}

	if err := database.DB.Delete(&blogpost).Error; err != nil {
		return c.Status(500).Render("errors/500", fiber.Map{"Error": "Error deleting post"})
	}

	return c.Redirect("/admin/dashboard")
}

// AdminNewProjectForm renders the add project form
func AdminNewProjectForm(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}
	return c.Render("admin/new_project", fiber.Map{
		"Title": "Add New Projects",
		"Admin": admin,
	})
}

// AdminCreateProject handles POST request to add a new project
func AdminCreateProject(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	project := models.Projects{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Link:        c.FormValue("link"),
		ImageURL:    c.FormValue("image"),
	}
	if err := database.DB.Create(&project).Error; err != nil {
		return c.Status(500).SendString("Error saving project")
	}

	return c.Redirect("/admin/dashboard")
}

// AdminDeleteProject deletes a project
func AdminDeleteProject(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	if err := database.DB.Delete(&models.Projects{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting project")
	}

	return c.Redirect("/admin/dashboard")
}


// create services functionality
func AdminCreateServices(c *fiber.Ctx) error {
	// fetch and check sessiona to ensure its the admin 
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}
	
	service := models.Services{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		//Link:        c.FormValue("link"),
		ImageURL:    c.FormValue("image"),
	}
	if err := database.DB.Create(&service).Error; err != nil {
		return c.Status(500).SendString("Error saving project")
	}
	
	return c.Status(200).SendString("Service Created Successfully")	
	
}

// AdminDeleteService deletes a service
func AdminDeleteService(c *fiber.Ctx) error {
	// fetch and check sessiona to ensure its the admin 
	sess, _ := config.Store.Get(c)
	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}
	
	// fetch service ID from session
	serviceid := c.Params("id")
	if serviceid == "" {
		return c.Status(404).SendString("Invalid Service ID")
	}

	// Delete service from database
	if err := database.DB.Delete(&models.Services{}, serviceid).Error; err != nil {
		return c.Status(500).SendString("Error deleting project")
	}
	
	return c.Status(200).SendString("Service deleted Successfully")
}