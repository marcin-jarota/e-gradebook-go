package main

import (
	"context"
	"e-student/internal/adapters/db"
	"e-student/internal/adapters/storage"
	"e-student/internal/adapters/transport"
	core "e-student/internal/app"
	"e-student/internal/auth"
	"e-student/internal/middleware"
	"e-student/internal/student"
	"e-student/internal/subject"
	"e-student/internal/user"
)

func main() {
	cfg := core.NewConfig()

	conn, closeDbConn := db.NewGormDB(cfg)
	server := transport.NewFiberHttpServer(cfg)
	storage, closeRedis := storage.NewRedisStorage("session", cfg.RedisAddr, context.Background())

	defer closeRedis()
	defer closeDbConn()
	defer server.App.Shutdown()

	// repositories
	userRepo := user.NewGormUserRepository(conn)
	studentRepo := student.NewGormStudentRepository(conn)
	subjectRepo := subject.NewGormSubjectRepository(conn)

	// services
	authService := auth.NewAuthService(userRepo, storage, cfg)
	studentService := student.NewStudentService(studentRepo)
	subjectService := subject.NewSubjectService(subjectRepo)
	userService := user.NewUserService(userRepo, storage)

	// middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// handlers
	studentHandler := student.NewStudentHandler(studentService)
	authHandler := auth.NewAuthHandler(authService)
	subjecthandler := subject.NewSubjectHandler(subjectService)
	userHandler := user.NewUserHandler(userService)

	// bind routing
	studentHandler.BindRouting(server.App, authMiddleware)
	authHandler.BindRouting(server.App, authMiddleware)
	subjecthandler.BindRouting(server.App, authMiddleware)
	userHandler.BindRouting(server.App, authMiddleware)

	server.Listen()
}
