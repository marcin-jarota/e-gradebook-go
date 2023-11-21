package classgroup

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type ClassGroupHandler struct {
	transport.Handler
	classGroupService ports.ClassGroupService
	studentsService   ports.StudentService
}

func NewClassGroupHandler(classGroupService ports.ClassGroupService, studentsService ports.StudentService) *ClassGroupHandler {
	return &ClassGroupHandler{
		classGroupService: classGroupService,
		studentsService:   studentsService,
	}
}

func (h *ClassGroupHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/class-groups", auth.IsAuthenticatedByHeader())
	r.Get("/", h.GetAll)
	r.Post("/", auth.IsAdmin(), h.Create)
	r.Get("/:classGroupID/students", auth.IsAdmin(), h.ListStudents)
	r.Post("/:classGroupID/students", auth.IsAdmin(), h.AddStudentToClassGroup)
}

func (h *ClassGroupHandler) ListStudents(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	output, err := h.studentsService.GetAllByClassGroup(classGroupID)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, output)
}

func (h *ClassGroupHandler) AddStudentToClassGroup(c *fiber.Ctx) error {
	var p ports.AddStudentToClassGroupPayload
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	setClassGroupData := ports.SetClassGroupPayload{
		StudentID:    p.StudentID,
		ClassGroupID: classGroupID,
	}

	if err := h.studentsService.SetClassGroup(setClassGroupData); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, nil)
}

func (h *ClassGroupHandler) Create(c *fiber.Ctx) error {
	var p ports.AddClassGroupInput

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := h.classGroupService.AddClassGroup(p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, fiber.Map{
		"success": true,
	})
}

func (h *ClassGroupHandler) GetAll(c *fiber.Ctx) error {
	classGroups, err := h.classGroupService.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, classGroups)
}
