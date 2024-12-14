package main

import (
	handler "github.com/j0kernotathome/Go-password-manager/UserService/internal/transport"
)

func main() {

	app := handler.CreateServer()
	app.Listen(":8080")

}
