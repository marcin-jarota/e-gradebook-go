package transport

import (
	"e-student/internal/app"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberHttpServer struct {
	cfg *app.Config
	App *fiber.App
}

func NewFiberHttpServer(cfg *app.Config) *FiberHttpServer {
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		Next:             nil,
	}))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3001,http://localhost:5173,http://127.0.0.1:5173",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	return &FiberHttpServer{
		App: app,
		cfg: cfg,
	}
}

func (s *FiberHttpServer) Listen() {
	log.Printf("Listening on port http://localhost:%s", s.cfg.Port)

	s.App.Static("/static", "./web/dist")
	s.App.Listen(fmt.Sprintf(":%s", s.cfg.Port))
}
