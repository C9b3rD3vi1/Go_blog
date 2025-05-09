package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/middleware"

)


// 
func SetupRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Get("/login", handlers.AdminAuthHandler)
	admin.Post("/login", handlers.AdminAuthHandler)


	admin.Use(middleware.RequireAdminAuth) // protect routes below

	admin.Get("/dashboard", handlers.AdminDashboard)
	admin.Get("/posts", handlers.AdminPostList)
	admin.Get("/posts/new", handlers.AdminPostForm)
	admin.Post("/posts", handlers.AdminCreatePost)
	admin.Get("/posts/edit/:id", handlers.AdminEditPostForm)
	admin.Post("/posts/update/:id", handlers.AdminUpdatePost)
	admin.Post("/posts/delete/:id", handlers.AdminDeletePost)
}
