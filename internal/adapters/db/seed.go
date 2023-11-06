package db

import (
	"e-student/internal/app/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB) {
	admin := &domain.User{
		Name:     "Marcin",
		Surname:  "Testowy",
		Password: "devpass",
		Email:    "admin@e-student.com",
		Role:     domain.AdminRole,
		Active:   true,
	}

	hash, err := admin.GeneratePassword(admin.Password)

	if err != nil {
		panic(err)
	}

	admin.Password = hash

	if err := db.First(&admin, "email = ?", "admin@e-gradebook.com").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[SEED]: ADMIN USER CREATED")
			db.Create(&admin)
		} else {
			panic(err)
		}
	}
}

func SeedSubject(db *gorm.DB) {
	name := "Wychowanie Fizyczne"
	subject := &domain.Subject{
		Name: name,
	}

	if err := db.First(&subject, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&subject)
			log.Println("[SEED]: SUBJECT ADDED")
		} else {
			panic(err)
		}
	}
}
