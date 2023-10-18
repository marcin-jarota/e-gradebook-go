package main

import (
	"context"
	"e-student/internal/adapters/storage"
	"e-student/internal/auth"
	"e-student/internal/common"
	"e-student/internal/middleware"
	"e-student/internal/student"
	"e-student/internal/user"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// gormLogger "gorm.io/gorm/logger"
)

func main() {
	cfg := common.NewConfig()

	// newLogger := gormLogger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	gormLogger.Config{
	// 		SlowThreshold:             time.Second,     // Slow SQL threshold
	// 		LogLevel:                  gormLogger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      false,           // Don't include params in the SQL log
	// 		Colorful:                  true,            // Disable color
	// 	},
	// )

	conn, err := gorm.Open(postgres.Open(cfg.Dsn)) // &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}

	storage := storage.NewRedisStorage("session", cfg.RedisAddr, context.Background())

	// repositories
	userRepo := user.NewGormUserRepository(conn)
	studentRepo := student.NewGormStudentRepository(conn)

	// services
	authService := auth.NewAuthService(userRepo, storage, cfg)
	studentService := student.NewStudentService(studentRepo)

	// middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService)

	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		Next:             nil,
	}))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3001,http://localhost:5173",
	}))

	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// bind routing
	student.NewStudentHandler(studentService, authMiddleware).BindRouting(app)
	user.NewUserHandler(authService).BindRouting(app)

	log.Printf("Listening on port http://localhost:%s", cfg.Port)

	app.Static("/static", "./web/dist")
	app.Listen(fmt.Sprintf(":%s", cfg.Port))
}
