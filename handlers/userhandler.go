package handlers

import "github.com/gofiber/fiber/v2"

func HomePageHandler(c *fiber.Ctx) error {
	// extract the user from the session
	//user := c.Locals("user").(*models.User)

	return c.Render("pages/index", fiber.Map{
		//	"User": user,
	})

}

func UserRegisterHandlerForm(c *fiber.Ctx) error {
	return c.Render("pages/register", fiber.Map{})

}

func UserLoginHandlerForm(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{})

}

func UserContactHandlerForm(c *fiber.Ctx) error {
	return c.Render("pages/contact", fiber.Map{})

}

func AboutUsHandler(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{})

}
