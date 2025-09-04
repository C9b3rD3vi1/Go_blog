package utils

import (
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

// Slugify converts text into a URL-friendly slug
func Slugify(s string) string {
	// Lowercase
	slug := strings.ToLower(s)
	// Replace non-alphanumeric with hyphen
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	// Trim hyphens
	slug = strings.Trim(slug, "-")
	return slug
}

// UniqueSlug ensures slug uniqueness for a given table
func UniqueSlug(db *gorm.DB, table string, title string) string {
	base := Slugify(title)
	slug := base
	count := 1

	var exists bool
	for {
		// Check if slug already exists in given table
		var result int64
		db.Table(table).Where("slug = ?", slug).Count(&result)

		if result == 0 {
			break // unique slug found
		}

		// Append -1, -2, etc.
		slug = fmt.Sprintf("%s-%d", base, count)
		count++
		exists = true
	}

	if exists {
		return slug
	}
	return base
}
