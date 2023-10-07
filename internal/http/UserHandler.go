package http

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	auth ports.AuthService
}

func NewUserHandler(auth ports.AuthService) *UserHandler {
	return &UserHandler{
		auth: auth,
	}
}

func (u *UserHandler) BindRouting(app *fiber.App) *UserHandler {
	app.Get("/login", u.GetLogin)
	app.Post("/login", u.PostLogin)
	app.Get("/logout", u.AuthMiddleware(), u.Logout)
	app.Get("/start", u.AuthMiddleware(), u.Start)

	return u
}

func (u *UserHandler) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isLoggedIn, user := u.auth.IsLoggedIn(c.Cookies("token"))

		if isLoggedIn {
			fmt.Println(user.Name)

			c.Locals("user", user)
			return c.Next()
		}

		return c.Redirect("/login")
	}
}

func (u *UserHandler) GetLogin(c *fiber.Ctx) error {
	isLoggedIn, _ := u.auth.IsLoggedIn(c.Cookies("token"))

	if isLoggedIn {
		return c.Redirect("/start")
	}

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

	return c.Redirect("/start", fiber.StatusFound)
}

func (u *UserHandler) Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.SessionUser)
	err := u.auth.Logout(int(user.ID))
	if err != nil {
		panic(err)
	}

	return c.Render("pages/login", fiber.Map{"Error": "Logged out!"}, "layouts/main")

}

func (u *UserHandler) Start(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.SessionUser)

	return c.Render("pages/start", fiber.Map{
		"User": user,
	}, "layouts/main")
}
