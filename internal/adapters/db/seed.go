package db

import (
	"e-student/internal/app/domain"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

var teacher = domain.User{
	Name:     "Adam",
	Surname:  "Kowalski",
	Password: "devpass",
	Email:    "teacher@e-student.com",
	Role:     domain.TeacherRole,
	Active:   true,
}

var admin = domain.User{
	Name:     "Marcin",
	Surname:  "Testowy",
	Password: "devpass",
	Email:    "admin@e-student.com",
	Role:     domain.AdminRole,
	Active:   true,
}

var student = domain.User{
	Name:     "Maciej",
	Surname:  "Szkolny",
	Password: "devpass",
	Email:    "student@e-student.com",
	Role:     domain.StudentRole,
	Active:   true,
}

func SeedAdminUser(db *gorm.DB) {

	hash, err := admin.GeneratePassword(admin.Password)

	if err != nil {
		panic(err)
	}

	admin.Password = hash

	if err := db.First(&admin, "email = ?", admin.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[SEED]: Creating admin user")
			db.Create(&admin)
		} else {
			panic(err)
		}
	}
}

func SeedStudentUser(db *gorm.DB) {
	hash, err := student.GeneratePassword(student.Password)

	if err != nil {
		panic(err)
	}

	var group domain.ClassGroup

	err = db.First(&group, "name = ?", "1 mat-fiz").Error

	if err != nil {
		panic(err)
	}

	student.Password = hash

	if err := db.Model(&domain.User{}).First(&student, "email = ?", student.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(fmt.Sprintf("[SEED]: Creating student with class group: %s", group.Name))
			db.Create(&domain.Student{
				User:       student,
				ClassGroup: group,
			})
		} else {
			panic(err)
		}
	}
}

func SeedTeacherUser(db *gorm.DB) {
	hash, err := teacher.GeneratePassword(teacher.Password)

	if err != nil {
		panic(err)
	}

	teacher.Password = hash

	if err := db.Model(&domain.User{}).First(&teacher, "email = ?", teacher.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			var subject domain.Subject
			var newTeacher domain.Teacher

			newTeacher.User = teacher

			err := db.FirstOrCreate(&subject, domain.Subject{Name: "Wychowanie Fizyczne"}).Error

			if err == nil {
				newTeacher.Subjects = append(newTeacher.Subjects, subject)
			} else {
				log.Println(err)
			}

			log.Println(fmt.Sprintf("[SEED]: Creating teacher user with subject assigned %s", subject.Name))

			db.Create(&newTeacher)
		} else {
			panic(err)
		}
	}
}

func SeedClassGroup(db *gorm.DB) {
	name := "1 mat-fiz"

	classGroup := &domain.ClassGroup{
		Name: name,
		// Students: []domain.Student{},
	}

	if err := db.First(&classGroup, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&classGroup)
			log.Println("[SEED]: CLASS GROUP ADDED")
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

	var myClass domain.ClassGroup
	err := db.FirstOrCreate(&myClass, domain.ClassGroup{Name: "1 mat-fiz"}).Error

	if err == nil {
		subject.ClassGroups = append(subject.ClassGroups, myClass)
	} else {
		log.Panic(err)
	}

	if err := db.First(&subject, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&subject)
			log.Println("[SEED]: Creating subject")
		} else {
			panic(err)
		}
	}
}

func SeedMarks(db *gorm.DB) {
	var s domain.Student
	var t domain.Teacher
	var sub domain.Subject

	if err := db.Find(&s, "id = ?", 1).Error; err != nil {
		panic(err)
	}

	if err := db.Find(&t, "id = ?", 1).Error; err != nil {
		panic(err)
	}

	if err := db.Find(&sub, "id = ?", 1).Error; err != nil {
		panic(err)
	}

	mark := domain.Mark{
		Value:   3.5,
		Subject: sub,
		Teacher: t,
		Student: s,
	}

	db.Create(&mark)
	log.Println(fmt.Sprintf("[SEED]: Creating mark %f for subject %s", mark.Value, mark.Subject.Name))
}
