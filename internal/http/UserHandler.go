package http

import (
	"e-student/internal/app/ports"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	auth ports.AuthService
}

func NewUserHandler(userUscase ports.AuthService) *UserHandler {
	return &UserHandler{
		auth: userUscase,
	}
}

func (u *UserHandler) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isLoggedIn, user := u.auth.IsLoggedIn(c.Cookies("token"))

		if isLoggedIn {
			fmt.Println(user.Name)

			c.Locals("User", user)
			return c.Next()
		}

		return c.Redirect("/")
	}
}

func (u *UserHandler) GetLogin(c *fiber.Ctx) error {
	return c.Render("pages/login", nil, "layouts/main")
}

func (u *UserHandler) PostLogin(c *fiber.Ctx) error {
	token, err := u.auth.Login(c.FormValue("email"), c.FormValue("password"))

	if err != nil {
		return c.Render("pages/login", fiber.Map{
			"Error": err.Error(),
		}, "layouts/main")
	}

	c.Cookie(&fiber.Cookie{Name: "token", Value: token})

	return c.Redirect("/app/home", fiber.StatusFound)
}
