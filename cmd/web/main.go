package main

import (
	"context"
	"e-student/internal/adapters/repository"
	"e-student/internal/adapters/service"
	"e-student/internal/adapters/storage"
	"e-student/internal/common"
	transport "e-student/internal/http"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := common.NewConfig()

	conn, err := gorm.Open(postgres.Open(cfg.Dsn), nil)

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewGormUserRepository(conn)
	markRepo := repository.NewGormMarkRepository(conn)
	studentRepo := repository.NewGormStudentRepository(conn)
	// subjectRepo := repository.NewGormSubjectRepository(conn)

	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{Views: engine, ViewsLayout: "layouts/main"})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		Next:             nil,
	}))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3001,http://localhost:5173",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	storage := storage.NewRedisStorage("session", cfg.RedisAddr, context.Background())
	authService := service.NewAuthService(userRepo, storage, cfg)
	studentSrvc := service.NewStudentService(studentRepo, markRepo)

	transport.NewStudentHandler(studentSrvc).BindRouting(app)
	transport.NewUserHandler(authService).BindRouting(app)

	log.Printf("Listening on port http://localhost:%s", cfg.Port)

	app.Static("/static", "./web/dist")
	app.Listen(fmt.Sprintf(":%s", cfg.Port))
	// http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)

}
