package main

import (
	"context"
	"e-student/internal/adapters/db"
	"e-student/internal/adapters/storage"
	"e-student/internal/adapters/transport"
	core "e-student/internal/app"
	"e-student/internal/app/ports"
	"e-student/internal/auth"
	classgroup "e-student/internal/class_group"
	"e-student/internal/lesson"
	"e-student/internal/mark"
	"e-student/internal/middleware"
	"e-student/internal/notification"
	"e-student/internal/student"
	"e-student/internal/subject"
	"e-student/internal/teacher"
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
	classgroupRepo := classgroup.NewClassGroupRepository(conn)
	markRepo := mark.NewGormMarkRepository(conn)
	teacherRepo := teacher.NewTeacherRepository(conn)
	lessonRepo := lesson.NewLessonRepository(conn)
	notificationRepo := notification.NewGormNotificationRepository(conn)

	// services
	authService := auth.NewAuthService(userRepo, storage, cfg)
	markService := mark.NewMarkService(markRepo)
	teacherService := teacher.NewTeacherService(teacherRepo)
	studentService := student.NewStudentService(studentRepo, markService)
	subjectService := subject.NewSubjectService(subjectRepo)
	userService := user.NewUserService(userRepo, storage)
	classgroupService := classgroup.NewClassGroupService(classgroupRepo)
	lessonService := lesson.NewLessonService(lessonRepo)
	notificationService := notification.NewNotificationService(
		map[string]ports.NotificationStrategy{
			"in-app": notification.NewInAppNotification(notificationRepo),
		},
		notificationRepo,
	)

	// middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// handlers
	studentHandler := student.NewStudentHandler(studentService, markService)
	authHandler := auth.NewAuthHandler(authService)
	subjecthandler := subject.NewSubjectHandler(subjectService)
	userHandler := user.NewUserHandler(userService, studentService, teacherService, authService, cfg)
	classgroupHandler := classgroup.NewClassGroupHandler(classgroupService, studentService, markService, lessonService)
	markHandler := mark.NewMarkHandler(markService, studentService, subjectService, notificationService)
	teacherHandler := teacher.NewTeacherHandler(teacherService)
	lessonHandler := lesson.NewLessonHandler(lessonService)
	notificationHandler := notification.NewNotificationHandler(notificationService)

	// bind routing
	studentHandler.BindRouting(server.App, authMiddleware)
	authHandler.BindRouting(server.App, authMiddleware)
	subjecthandler.BindRouting(server.App, authMiddleware)
	userHandler.BindRouting(server.App, authMiddleware)
	classgroupHandler.BindRouting(server.App, authMiddleware)
	markHandler.BindRouting(server.App, authMiddleware)
	teacherHandler.BindRouting(server.App, authMiddleware)
	lessonHandler.BindRouting(server.App, authMiddleware)
	notificationHandler.BindRouting(server.App, authMiddleware)

	server.Listen()
}
