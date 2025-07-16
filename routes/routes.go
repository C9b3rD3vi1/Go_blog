package routes

import (
	"github.com/C9b3rD3vi1/Go_blog/auth"
	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Get("/login", auth.AdminAuthHandler)
	admin.Post("/login", auth.AdminAuthHandler)

	admin.Use(middleware.RequireAdminAuth) // protect routes below

	admin.Get("/dashboard", handlers.AdminDashboard)
	admin.Get("/posts", handlers.AdminPostList)
	//admin.Get("/posts/new", handlers.AdminPostForm)
	//	admin.Post("/posts", handlers.AdminCreatePost)
	admin.Get("/posts/edit/:id", handlers.AdminEditPostForm)
	admin.Post("/posts/update/:id", handlers.AdminUpdatePost)
	admin.Post("/posts/delete/:id", handlers.AdminDeletePost)
}
