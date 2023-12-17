package main

import (
	"e-student/internal/adapters/db"
	"e-student/internal/app"
	"e-student/internal/app/domain"
	"fmt"
	"log"
)

func main() {
	cfg := app.NewConfig()
	conn, closeConn := db.NewGormDB(cfg)

	defer closeConn()

	db.AskForCleanDatabase(conn)

	userRoleDrop := `DROP TYPE IF EXISTS user_role`
	userRoleCreate := fmt.Sprintf(`CREATE TYPE user_role AS ENUM ('%s', '%s', '%s');`, domain.AdminRole, domain.StudentRole, domain.TeacherRole)

	if err := conn.Exec(userRoleDrop).Error; err != nil {
		log.Println(err)
	} else {
		if err := conn.Exec(userRoleCreate).Error; err != nil {
			fmt.Printf("Error creating enum type: %v\n", err)
		} else {
			fmt.Printf("[SEED]: user_role enum created!")
		}
	}

	if err := conn.AutoMigrate(&domain.User{}); err != nil {
		log.Panic("Could not boot User table: ", err)
	}

	if err := conn.AutoMigrate(&domain.Student{}); err != nil {
		log.Panic("Could not boot Student table: ", err)
	}

	if err := conn.AutoMigrate(&domain.Subject{}); err != nil {
		log.Panic("Could migrate Subject table: ", err)
	}

	if err := conn.AutoMigrate(&domain.Mark{}); err != nil {
		log.Panic("Could migrate Mark table: ", err)
	}

	if err := conn.AutoMigrate(&domain.ClassGroup{}); err != nil {
		log.Panic("Could migrate ClassGroup table: ", err)
	}

	if err := conn.AutoMigrate(&domain.Teacher{}); err != nil {
		log.Panic("Could migrate Teacher table: ", err)
	}

	if err := conn.AutoMigrate(&domain.SubjectTeacherClass{}); err != nil {
		log.Panic(`Could not migrate SubjectTeacherClass: `, err)
	}

	if err := conn.AutoMigrate(&domain.Lesson{}); err != nil {
		log.Panic("Could not migrate Lesson", err)
	}

	if err := conn.AutoMigrate(&domain.Notification{}); err != nil {
		log.Panic("Could not migrate Notificationj", err)
	}

	if err := conn.AutoMigrate(&domain.SchoolYear{}); err != nil {
		log.Panic("Could not migrate SchoolYear", err)
	}

	db.LoadSQLFile("open_school_year.sql", conn)
	db.LoadSQLFile("trigger_set_school_year.sql", conn)

	db.SeedAdminUser(conn)
	db.SeedTeacherUser(conn)
	db.SeedSchoolYear(conn)
	db.SeedSubject(conn)
	db.SeedClassGroup(conn)
	db.SeedStudentUser(conn)
	db.SeedMarks(conn)
	db.SeedSubjectTeacherClassGroup(conn)
	db.SeedLessons(conn)
}
