package auth

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	transport.Handler
	service ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) BindRouting(app *fiber.App) *AuthHandler {
	app.Get("/login", h.GetLogin)
	app.Post("/login", h.PostLoginVue)
	app.Get("/logout", h.Logout)

	return h
}

func (h *AuthHandler) GetLogin(c *fiber.Ctx) error {
	isLoggedIn, _ := h.service.IsLoggedIn(c.Cookies("token"))

	if isLoggedIn {
		return c.Redirect("/start")
	}

	return c.Render("pages/login", nil, "layouts/login")
}

func (h *AuthHandler) PostLogin(c *fiber.Ctx) error {
	token, err := h.service.Login(c.FormValue("email"), c.FormValue("password"))

	if err != nil {
		return c.Render("pages/login", fiber.Map{
			"Error": err.Error(),
		}, "layouts/main")
	}

	c.Cookie(&fiber.Cookie{Name: "token", Value: token})

	return c.Redirect("/start", fiber.StatusFound)
}

func (h *AuthHandler) PostLoginVue(c *fiber.Ctx) error {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Could not parse message",
		})
	}

	token, err := h.service.Login(payload.Email, payload.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.User)
	err := h.service.Logout(int(user.ID))
	if err != nil {
		panic(err)
	}

	return c.Render("pages/login", fiber.Map{"Error": "Logged out!"}, "layouts/main")

}
