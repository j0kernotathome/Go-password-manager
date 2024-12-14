package handler

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	Models "github.com/j0kernotathome/Go-password-manager/UserService/internal/models/Users"
	"github.com/j0kernotathome/Go-password-manager/UserService/internal/singleton"
)

func getUser(ctx *fiber.Ctx) error {
	login := ctx.Params("login")
	var u Models.User
	u.GetUserByLogin(login)
	if u.Login == "" || u.Email == "" || u.Password == "" {
		ctx.Write([]byte("Invalid values"))
		return ctx.SendStatus(400)
	}
	out, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}
	ctx.Write(out)
	return ctx.SendStatus(200)
}
func createUser(ctx *fiber.Ctx) error {
	db := singleton.ConnectToDb()
	var body Models.User
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Println(err)
	}
	db.Model(&Models.User{}).Count(&body.Id)
	err = Models.CreateUserInDb(body)
	if err != nil {
		log.Println(err)
		ctx.Write([]byte(err.Error()))
		return ctx.SendStatus(400)

	}
	ctx.Cookie(&fiber.Cookie{
		Name:  "login",
		Value: body.Login,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:  "password",
		Value: body.Password,
	})
	return ctx.SendStatus(200)
}
