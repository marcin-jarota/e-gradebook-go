package mark

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type markHandler struct {
	transport.Handler
	markService ports.MarkService
}

func NewMarkHandler(markService ports.MarkService) *markHandler {
	return &markHandler{
		markService: markService,
	}
}

func (h *markHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {

}
