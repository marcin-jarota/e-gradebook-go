package user

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	transport.Handler
	service        ports.UserService
	studentService ports.StudentService
	teacherService ports.TeacherService
	auth           ports.AuthService
	cfg            *app.Config
}

func NewUserHandler(service ports.UserService, studentService ports.StudentService, teacherService ports.TeacherService, auth ports.AuthService, cfg *app.Config) *UserHandler {
	return &UserHandler{
		service:        service,
		auth:           auth,
		studentService: studentService,
		teacherService: teacherService,
		cfg:            cfg,
	}
}

func (h *UserHandler) BindRouting(app *fiber.App, auth *middleware.AuthMiddleware) {
	r := app.Group("/user", auth.IsAuthenticatedByHeader())

	r.Get("/list", auth.IsAdmin(), h.GetAll)
	r.Get("/activate/:id", auth.IsAdmin(), h.ActivateUser)
	r.Get("/deactivate/:id", auth.IsAdmin(), h.DeactivateUser)
	r.Get("/destroy-session/:id", auth.IsAdmin(), h.DestroyUserSession)
	r.Post("/create/:role", auth.IsAdmin(), h.PostAddUser)
	r.Get("/:id/student-data", auth.UserIs("student"), h.GetStudentByUserID)
	r.Get("/:id/teacher-data", auth.UserIs("teacher"), h.GetTeacherByUserID)

	app.Post("/setup-password", h.SetupPassword)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, users)
}

func (h *UserHandler) SetupPassword(c *fiber.Ctx) error {
	var payload ports.SetupPasswordPayload
	params := c.Queries()

	if err := c.BodyParser(&payload); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	email, userID, isValid := h.auth.IsGenerateTokenValid(params["token"])

	fmt.Println(email, userID, isValid)
	if !isValid {
		return h.JSONError(c, errors.New("token invalid"), fiber.StatusBadRequest)
	}

	if err := h.service.SetupPassword(email, payload.Password, payload.PasswordConfirm); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	if err := h.service.Activate(userID); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, fiber.Map{
		"ok": true,
	})
}

func (h *UserHandler) ActivateUser(c *fiber.Ctx) error {
	intId, err := strconv.Atoi(c.Params("id", "0"))

	if intId == 0 || err != nil {
		return h.JSONError(c, errors.New("incorrect user id"), fiber.StatusBadRequest)
	}

	err = h.service.Activate(uint(intId))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, fiber.Map{
		"ok": true,
	})
}

func (h *UserHandler) DeactivateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id", "0"))

	if id == 0 || err != nil {
		return h.JSONError(c, errors.New("incorrect user id"), fiber.StatusBadRequest)
	}

	err = h.service.Deactivate(uint(id))

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

func (h *UserHandler) PostAddUser(c *fiber.Ctx) error {
	role := c.Params("role", "student")
	var p ports.UserCreatePayload

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	var err error

	if role == "student" {
		err = h.studentService.AddStudent(p)
	}

	if role == "admin" {
		err = h.service.AddAdmin(p)
	}

	if role == "teacher" {
		err = h.teacherService.AddTeacher(p)
	}

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	token, err := h.auth.GeneratePasswordToken(p.Email, 15*time.Minute)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, map[string]any{
		"activationLink": fmt.Sprintf("%s/setup-password?token=%s&email=%s", h.cfg.BaseUrlFront, token, p.Email),
	})
}

func (h *UserHandler) GetStudentByUserID(c *fiber.Ctx) error {
	id, err := h.ParseIntParam(c.Params("id", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	student, err := h.studentService.GetByUserID(id)

	if err != nil {
		h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, student)
}

func (h *UserHandler) GetTeacherByUserID(c *fiber.Ctx) error {
	id, err := h.ParseIntParam(c.Params("id", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	teacher, err := h.teacherService.GetTeacherByUserID(id)

	if err != nil {
		h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, teacher)
}
