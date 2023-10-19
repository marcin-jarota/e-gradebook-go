package main

import (
	"e-student/internal/adapters/db"
	"e-student/internal/app"
	"e-student/internal/app/domain"
	"log"
)

func main() {
	cfg := app.NewConfig()
	conn, closeConn := db.NewGormDB(cfg)

	defer closeConn()

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

	db.SeedAdminUser(conn)
	db.SeedSubject(conn)
}
