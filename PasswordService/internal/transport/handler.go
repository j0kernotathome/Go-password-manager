package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func CreateServer() *fiber.App {
	app := fiber.New()
	app.Post("/add_password", addPassword)
	app.Get("/get_passwords", getPasswords)
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "KslZBuSPIROwMgGvtuWIUjKspDSDtZvF",
	}))
	return app
}
