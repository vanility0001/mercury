package main

import (
	b64 "encoding/base64"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Response struct {
	// http error code
	Success int `json:"success"`

	// response message
	Message string `json:"message"`

	// errors (if there are any)
	Errors []string `json:"errors"`
}

type Account struct {
	// account username
	Username string `json:"username"`

	// account password
	Password string `json:"password"`

	// account email
	Email string `json:"email"`
}

func main() {
	app := fiber.New()
	var nuts string = "/"

	app.Get(nuts, func(ctx *fiber.Ctx) error {
		ctx.Status(200)
		return ctx.JSON(&Response{
			Success: 200,
			Message: "ur gay",
			Errors:  nil,
		})
	})

	app.Post("api/register", func(ctx *fiber.Ctx) error {
		body := new(Account)
		if err := ctx.BodyParser(body); err != nil {
			ctx.Status(403)
			return ctx.JSON(&Response{
				Success: 403,
				Message: "Unknown error",
				Errors:  nil,
			})
		}

		ctx.Status(200)
		return ctx.JSON(&Account{
			Username: body.Username,
			Password: b64.StdEncoding.EncodeToString([]byte("a")),
			Email:    body.Email,
		})
	})

	/*app.Get("api/register", func(ctx *fiber.Ctx) error {
		return ctx.SendString("balls")
	})*/

	log.Fatal(app.Listen(":8080"))
}
