package user

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	transport.Handler
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) BindRouting(app *fiber.App, auth *middleware.AuthMiddleware) {
	r := app.Group("/user", auth.IsAuthenticatedByHeader(), auth.IsAdmin())

	r.Get("/list", h.GetAll)
	r.Get("/activate/:id", h.ActivateUser)
	r.Get("/destroy-session/:id", h.DestroyUserSession)
	r.Post("/create", h.PostAddAdmin)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, users)
}

func (h *UserHandler) ActivateUser(c *fiber.Ctx) error {
	intId, err := strconv.Atoi(c.Params("id", "0"))

	if intId == 0 {
		return h.JSONError(c, errors.New("incorrect user id"), fiber.StatusBadRequest)
	}

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	err = h.service.Activate(uint(intId))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, fiber.Map{
		"ok": true,
	})
}

func (h *UserHandler) DestroyUserSession(c *fiber.Ctx) error {
	intId, err := strconv.Atoi(c.Params("id", "0"))

	if intId == 0 {
		return h.JSONError(c, errors.New("incorrect user id"), fiber.StatusBadRequest)
	}

	err = h.service.DestroySession(uint(intId))

	if err != nil {

		return h.JSONError(c, err, fiber.StatusInternalServerError)

	}

	return h.JSON(c, fiber.Map{
		"ok": true,
	})
}

func (h *UserHandler) PostAddAdmin(c *fiber.Ctx) error {
	var p ports.AdminCreatePayload

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	err := h.service.AddAdmin(&p)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, map[string]any{
		"ok": true,
	})
}
