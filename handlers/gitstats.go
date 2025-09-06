package handlers

import(
	"net/http"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)


func GitHubStatsHandler(c *fiber.Ctx) error {
	resp, err := http.Get("https://api.github.com/repos/C9b3rD3vi1/Go_blog/contributors")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "GitHub API error"})
	}
	defer resp.Body.Close()

	var contributors []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&contributors)

	return c.JSON(fiber.Map{
		"contributors": len(contributors),
	})
}
