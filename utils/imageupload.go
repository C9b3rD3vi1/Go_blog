package utils

import (
    "fmt"
    "time"
    "strconv"

    "github.com/gofiber/fiber/v2"
)

// UploadImage saves the uploaded file to /uploads and returns its URL
func UploadImage(c *fiber.Ctx, field string) (string, error) {
    // Get file from form field (e.g., "image")
    file, err := c.FormFile(field)
    if err != nil {
        return "", fmt.Errorf("no file uploaded in field: %s", field)
    }

    // Generate unique filename
    filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)

    // Save to uploads/ folder
    if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename)); err != nil {
        return "", fmt.Errorf("failed to save file: %v", err)
    }

    // Return the public URL
    return "/uploads/" + filename, nil
}


// ParseInt safely parses string to int with default fallback
func ParseInt(s string, defaultValue ...int) int {
    val, err := strconv.Atoi(s)
    if err != nil && len(defaultValue) > 0 {
        return defaultValue[0]
    }
    return val
}