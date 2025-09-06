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
    // Posts CRUD
	admin.Get("/posts", handlers.AdminPostList)               // List all posts
	admin.Get("/posts/new", handlers.AdminFetchTags)        // Fetch all tags and render them in the template
	admin.Get("/posts/new", handlers.AdminNewPostForm)        // Show create form
	admin.Post("/posts", handlers.AdminCreatePost)            // Handle create
	admin.Get("/posts/:slug", handlers.AdminViewPosts)        //View single post
	admin.Get("/posts/edit/:id", handlers.AdminEditPostsForm) // Show edit form
	admin.Post("/posts/update/:id", handlers.AdminUpdatePost) // Handle update
	admin.Post("/posts/delete/:id", handlers.AdminDeletePost) // Handle delete


    // Projects admin routes
    admin.Get("/projects", handlers.AdminProjectList)           // list all
    admin.Get("/projects/new", handlers.AdminNewProjectPage)  // LIST ALL TECH STACK FOR MULTI SELECT
    admin.Get("/projects/new", handlers.AdminNewProjectForm)    // show create form
    admin.Post("/projects/new", handlers.AdminCreateProject)    // handle create
    admin.Get("/projects/view/:slug", handlers.AdminViewProject) // view single project by slug
    admin.Get("/projects/edit/:id", handlers.AdminEditProjectForm) // show edit form
    admin.Post("/projects/edit/:id", handlers.AdminUpdateProject)  // handle update
    admin.Get("/projects/delete/:id", handlers.AdminDeleteProject) // delete


    
    // Services
    // Services Admin Routes
    admin.Get("/services", handlers.AdminServiceList)            // List all services
    admin.Get("/services/new", handlers.AdminNewServicePage)  // LIST ALL TECH STACK FOR MULTI SELECT
    admin.Get("/services/new", handlers.AdminNewServiceForm)    // Show form to create
    admin.Post("/services/new", handlers.AdminCreateServices)   // Handle create
    admin.Get("/services/edit/:id", handlers.AdminEditServiceForm)  // Show edit form
    admin.Post("/services/edit/:id", handlers.AdminUpdateService)   // Handle update
    admin.Get("/services/delete/:id", handlers.AdminDeleteService)  // Delete
    admin.Get("/services/:id", handlers.AdminViewService)       // View single service
    
    // Tech Stack Routes (admin)
    admin.Get("/techstacks", handlers.AdminTechStackList)
    admin.Get("/techstacks/new", handlers.AdminNewTechStackForm)
    admin.Post("/techstacks/new", handlers.AdminCreateTechStack)
    admin.Get("/techstacks/edit/:id", handlers.AdminEditTechStackForm)
    admin.Post("/techstacks/edit/:id", handlers.AdminUpdateTechStack)
    admin.Get("/techstacks/delete/:id", handlers.AdminDeleteTechStack)

    
    app.Get("/admin/tags",	handlers.AdminListTags)        // list all tags
    app.Post("/admin/tags", handlers.AdminCreateTag)      // create new tag
    app.Post("/admin/tags/delete/:id", handlers.AdminDeleteTag) // delete tag
}
