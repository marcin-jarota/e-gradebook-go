package classgroup

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

type classGroupHandler struct {
	transport.Handler
	classGroupService ports.ClassGroupService
	studentsService   ports.StudentService
	markService       ports.MarkService
	lessonService     ports.LessonService
}

func NewClassGroupHandler(classGroupService ports.ClassGroupService, studentsService ports.StudentService, markService ports.MarkService, lessonService ports.LessonService) *classGroupHandler {
	return &classGroupHandler{
		classGroupService: classGroupService,
		studentsService:   studentsService,
		markService:       markService,
		lessonService:     lessonService,
	}
}

func (h *classGroupHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/class-groups", auth.IsAuthenticatedByHeader())
	r.Get("/", h.GetAll)
	r.Post("/", auth.IsAdmin(), h.Create)
	r.Get("/:classGroupID", auth.UserIs("admin", "teacher"), h.Details)
	r.Delete("/:classGroupID", auth.UserIs("admin"), h.Delete)
	r.Get("/:classGroupID/students", auth.UserIs("admin", "teacher"), h.ListStudents)
	r.Get("/:classGroupID/marks", auth.UserIs("admin", "teacher"), h.ListMarks)
	r.Get("/:classGroupID/teachers", auth.UserIs("admin", "teacher"), h.ListTeachers)
	r.Get("/:classGroupID/teacher-subject", h.ListTeacherSubject)
	r.Get("/:classGroupID/lessons", h.ListLessons)
	r.Post("/:classGroupID/teachers", auth.IsAdmin(), h.AddTeacherToClassGroup)
	r.Post("/:classGroupID/students", auth.IsAdmin(), h.AddStudentToClassGroup)
	r.Post("/:classGroupID/subjects", auth.IsAdmin(), h.AssignTeacherWithSubject)
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

func (h *classGroupHandler) Delete(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	err = h.classGroupService.Delete(classGroupID)

	if err != nil {
		log.Println((err))
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, fiber.Map{"ok": true})
}

func (h *classGroupHandler) ListLessons(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	lessons, err := h.lessonService.GetByClassGroup(classGroupID)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, lessons)
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

func (h *classGroupHandler) ListTeachers(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	teachers, err := h.classGroupService.GetTeachers(classGroupID)

	if err != nil {

		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, teachers)
}

func (h *classGroupHandler) ListTeacherSubject(c *fiber.Ctx) error {
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	teachers, err := h.classGroupService.GetTeachersWithSubject(classGroupID)

	if err != nil {

		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, teachers)
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

func (h *classGroupHandler) AddTeacherToClassGroup(c *fiber.Ctx) error {
	var p ports.TeacherClassGroup
	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	p.ClassGroupID = classGroupID

	if err := h.classGroupService.AddTeacher(p); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, nil)
}

func (h *classGroupHandler) AssignTeacherWithSubject(c *fiber.Ctx) error {
	var p ports.TeacherSubjectClassgroupID

	classGroupID, err := h.ParseIntParam(c.Params("classGroupID", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	p.ClassGroupID = classGroupID

	if err := h.classGroupService.AddTeacherWithSubject(p); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, fiber.Map{"ok": true})
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
