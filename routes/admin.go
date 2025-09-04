package routes

import (
	"github.com/C9b3rD3vi1/Go_blog/auth"
	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupAdminRoutes sets up the admin routes for the application.
func SetupAdminRoutes(app *fiber.App) {
    // --- Public admin routes (NO middleware) ---
    app.Get("/admin/login", auth.AdminLoginForm)     // GET form
    app.Post("/admin/login", auth.AdminAuthHandler)    // POST form

    // --- Protected admin routes ---
    admin := app.Group("/admin", middleware.RequireAdminAuth)

    admin.Get("/logout", handlers.AdminLogout)
    admin.Get("/dashboard", handlers.AdminDashboard)
    admin.Get("/profile", handlers.AdminProfile)

    // User management
    admin.Get("/users", handlers.AdminUserList)
    admin.Get("/users/new", handlers.AdminCreateUser)
    admin.Post("/users/new", handlers.AdminCreateUser)

    // Posts
    admin.Get("/posts", handlers.AdminPostList)
    admin.Get("/posts/edit/:id", handlers.AdminEditPostForm)
    admin.Post("/posts/edit/:id", handlers.AdminUpdatePost)
    admin.Get("/posts/delete/:id", handlers.AdminDeletePost)

    // Projects admin routes
    //admin.Get("/projects", handlers.AdminProjectList)
    admin.Get("/projects/new", handlers.AdminNewProjectForm)
    admin.Post("/projects/new", handlers.AdminCreateProject)
    admin.Get("/projects/delete/:id", handlers.AdminDeleteProject)

    
    // Services
    // Services Admin Routes
    admin.Get("/services", handlers.AdminServiceList)            // List all services
    admin.Get("/services/new", handlers.AdminNewServiceForm)    // Show form to create
    admin.Post("/services/new", handlers.AdminCreateServices)   // Handle create
    admin.Get("/services/edit/:id", handlers.AdminEditServiceForm)  // Show edit form
    admin.Post("/services/edit/:id", handlers.AdminUpdateService)   // Handle update
    admin.Get("/services/delete/:id", handlers.AdminDeleteService)  // Delete
    admin.Get("/services/:id", handlers.AdminViewService)       // View single service

}
