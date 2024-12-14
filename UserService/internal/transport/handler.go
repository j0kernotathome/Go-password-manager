package handler

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func CreateServer() *fiber.App {
	app := fiber.New()
	app.Get("/get_user/:login", getUser)
	app.Post("/create_user", createUser)
	app.Get("/create_api_key", createApiKey)
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "KslZBuSPIROwMgGvtuWIUjKspDSDtZvF",
	}))
	return app
}

func createApiKey(ctx *fiber.Ctx) error {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	apiKey := make([]byte, 16)
	for i := 0; i < 16; i++ {
		apiKey[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	ctx.Write(apiKey)
	log.Println(string(apiKey))
	return ctx.SendStatus(200)
}
