package db

import "github.com/gofiber/fiber/v2/middleware/session"

var Store *session.Store

func InitStore() {
	session := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   true,
	})

	Store = session
}
