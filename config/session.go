package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

// InitSession initializes the session store with proper Fiber config
func InitSession() {
	Store = session.New(session.Config{
		KeyLookup:      "cookie:goblog", // Cookie name
		Expiration:     24 * time.Hour,  // Session expires in 24 hours
		CookiePath:     "/",             // Available on all routes
		CookieDomain:   "localhost",     // Change this in production
		CookieSecure:   false,           // Set true if using HTTPS
		CookieHTTPOnly: true,            // Prevents JS access to cookie
	})
}
