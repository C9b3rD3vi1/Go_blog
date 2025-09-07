package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)


func GitHubStatsHandler(c *fiber.Ctx) error {
	// Get repository details
	repoResp, err := http.Get("https://api.github.com/repos/C9b3rD3vi1/Go_blog")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "GitHub Repo API error"})
	}
	defer repoResp.Body.Close()

	var repoData map[string]interface{}
	if err := json.NewDecoder(repoResp.Body).Decode(&repoData); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to decode repo data"})
	}

	// Get contributors
	contribResp, err := http.Get("https://api.github.com/repos/C9b3rD3vi1/Go_blog/contributors")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "GitHub Contributors API error"})
	}
	defer contribResp.Body.Close()

	var contributors []map[string]interface{}
	if err := json.NewDecoder(contribResp.Body).Decode(&contributors); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to decode contributors"})
	}

	// Prepare response
	return c.JSON(fiber.Map{
		"stars":        repoData["stargazers_count"],
		"forks":        repoData["forks_count"],
		"open_issues":  repoData["open_issues_count"],
		"contributors": len(contributors),
	})
}


func GitHubUserStatsHandler(c *fiber.Ctx) error {
	// Replace with your GitHub username
	username := "C9b3rD3vi1"
	log.Printf("Fetching GitHub stats for user: %s", username)

	// --- Fetch user profile ---
	userResp, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		log.Printf("ERROR: failed to fetch user profile: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "GitHub User API error"})
	}
	defer userResp.Body.Close()

	log.Printf("GitHub user API response status: %s", userResp.Status)

	var userData map[string]interface{}
	if err := json.NewDecoder(userResp.Body).Decode(&userData); err != nil {
		log.Printf("ERROR: failed to decode user data: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to decode user data"})
	}
	log.Printf("Decoded user data: repos=%v followers=%v following=%v gists=%v",
		userData["public_repos"], userData["followers"], userData["following"], userData["public_gists"])

	// --- Fetch contributions SVG ---
	contriResp, err := http.Get("https://github.com/users/" + username + "/contributions")
	if err != nil {
		log.Printf("ERROR: failed to fetch contributions: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "GitHub Contributions API error"})
	}
	defer contriResp.Body.Close()

	log.Printf("GitHub contributions response status: %s", contriResp.Status)

	contributionsHTML, err := io.ReadAll(contriResp.Body)
if err != nil {
    log.Printf("ERROR: failed to read contributions page: %v", err)
    return c.Status(500).JSON(fiber.Map{"error": "Failed to read contributions page"})
}

svgRegex := regexp.MustCompile(`(?s)<svg.*?</svg>`)
svgMatches := svgRegex.FindAll(contributionsHTML, -1)

svg := ""
if len(svgMatches) > 0 {
    // pick the last <svg>, usually the contributions grid
    svg = string(svgMatches[len(svgMatches)-1])

    // 🎨 Replace default GitHub greens with indigo shades
    colorMap := map[string]string{
        "#ebedf0": "#f3f4f6", // very light gray → light background
        "#9be9a8": "#c7d2fe", // light indigo
        "#40c463": "#818cf8", // medium indigo
        "#30a14e": "#4f46e5", // deep indigo
        "#216e39": "#312e81", // darkest indigo
    }

    for old, new := range colorMap {
        svg = strings.ReplaceAll(svg, old, new)
    }
} else {
    log.Println("⚠️ Could not extract contributions SVG from page")
}

	if svg != "" {
		log.Printf("Extracted contributions SVG length: %d bytes", len(svg))
	} else {
		log.Println("⚠️ Could not extract SVG from contributions page")
	}

	// Prepare response
	response := fiber.Map{
		"public_repos":      userData["public_repos"], // number of repos
		"followers":         userData["followers"],    // followers count
		"following":         userData["following"],    // following count
		"gists":             userData["public_gists"], // public gists
		"contributions_svg": svg, // embed as string
	}

	log.Printf("Returning GitHub stats response")
	return c.JSON(response)
}
