package routes

import (
    "github.com/C9b3rD3vi1/Go_blog/handlers"
    "github.com/C9b3rD3vi1/Go_blog/middleware"
    "github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
    admin := app.Group("/admin")
    
    // Public routes
    //admin.Get("/login", handlers.AdminLogin)
    //admin.Post("/login", handlers.AdminLoginPost)
    
    // Protected admin routes
    admin.Use(middleware.RequireAdminAuth)
    admin.Get("/logout", handlers.AdminLogout)
    admin.Get("/dashboard", handlers.AdminDashboard)
    admin.Get("/profile", handlers.AdminProfile)
    
    // User management
    admin.Get("/users", handlers.AdminUserList)
    admin.Get("/users/new", handlers.AdminCreateUser)
    admin.Post("/users/new", handlers.AdminCreateUser)
    
    // Existing post and project routes
    admin.Get("/posts", handlers.AdminPostList)
    admin.Get("/posts/edit/:id", handlers.AdminEditPostForm)
    admin.Post("/posts/edit/:id", handlers.AdminUpdatePost)
    admin.Get("/posts/delete/:id", handlers.AdminDeletePost)
    
    admin.Get("/projects/new", handlers.AdminNewProjectForm)
    admin.Post("/projects/new", handlers.AdminCreateProject)
    admin.Get("/projects/delete/:id", handlers.AdminDeleteProject)
    
    admin.Post("/services/new", handlers.AdminCreateServices)
    admin.Get("/services/delete/:id", handlers.AdminDeleteService)
}