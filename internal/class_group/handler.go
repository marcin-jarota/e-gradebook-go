package classgroup

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type classGroupHandler struct {
	transport.Handler
	classGroupService ports.ClassGroupService
	studentsService   ports.StudentService
	markService       ports.MarkService
}

func NewClassGroupHandler(classGroupService ports.ClassGroupService, studentsService ports.StudentService, markService ports.MarkService) *classGroupHandler {
	return &classGroupHandler{
		classGroupService: classGroupService,
		studentsService:   studentsService,
		markService:       markService,
	}
}

func (h *classGroupHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/class-groups", auth.IsAuthenticatedByHeader())
	r.Get("/", h.GetAll)
	r.Post("/", auth.IsAdmin(), h.Create)
	r.Get("/:classGroupID", auth.IsAdmin(), h.Details)
	r.Get("/:classGroupID/students", auth.IsAdmin(), h.ListStudents)
	r.Get("/:classGroupID/marks", auth.UserIs("admin", "teacher"), h.ListMarks)
	r.Post("/:classGroupID/students", auth.IsAdmin(), h.AddStudentToClassGroup)
	r.Post("/:classGroupID/subjects", auth.IsAdmin(), h.AddSubjectToClassGroup)
}

func (h *classGroupHandler) Details(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	res, err := h.classGroupService.GetOneByID(classGroupID)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, res)
}

func (h *classGroupHandler) ListMarks(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	res, err := h.markService.GetByClassGroup(classGroupID)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, res)
}

func (h *classGroupHandler) ListStudents(c *fiber.Ctx) error {
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

func (h *classGroupHandler) AddStudentToClassGroup(c *fiber.Ctx) error {
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

func (h *classGroupHandler) AddSubjectToClassGroup(c *fiber.Ctx) error {
	var p ports.AddSubjectToClassGroupPayload
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	p.ClassGroupID = classGroupID

	if err := h.classGroupService.AddSubject(p); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, nil)
}

func (h *classGroupHandler) Create(c *fiber.Ctx) error {
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

func (h *classGroupHandler) GetAll(c *fiber.Ctx) error {
	classGroups, err := h.classGroupService.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, classGroups)
}
