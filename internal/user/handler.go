package user

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"e-student/internal/transport"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	transport.Handler
	auth ports.AuthService
}

func NewUserHandler(auth ports.AuthService) *UserHandler {
	return &UserHandler{
		auth: auth,
	}
}

func (u *UserHandler) BindRouting(app *fiber.App) *UserHandler {
	app.Get("/login", u.GetLogin)
	app.Post("/login", u.PostLoginVue)
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

	return c.Render("pages/login", nil, "layouts/login")
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

func (u *UserHandler) PostLoginVue(c *fiber.Ctx) error {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Could not parse message",
		})
	}

	token, err := u.auth.Login(payload.Email, payload.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (u *UserHandler) Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.User)
	err := u.auth.Logout(int(user.ID))
	if err != nil {
		panic(err)
	}

	return c.Render("pages/login", fiber.Map{"Error": "Logged out!"}, "layouts/main")

}

func (u *UserHandler) Start(c *fiber.Ctx) error {
	return u.RenderWithGlobals(c, "pages/start", nil)
}
