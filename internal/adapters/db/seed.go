package db

import (
	"e-student/internal/app/domain"
	"errors"
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

	err = db.FirstOrCreate(&group, "name = ?", "1 mat-fiz").Error

	if err != nil {
		panic(err)
	}

	student.Password = hash

	if err := db.Model(&domain.User{}).First(&student, "email = ?", student.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[SEED]: Creating student with class group: %s", group.Name)
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

			var newTeacher domain.Teacher

			newTeacher.User = teacher

			db.Create(&newTeacher)
		} else {
			panic(err)
		}
	}
}

func SeedClassGroup(db *gorm.DB) {
	name := "1 mat-fiz"
	var classGroup domain.ClassGroup

	if err := db.FirstOrCreate(&classGroup, "name = ?", name).Error; err != nil {
		log.Panic(err)
	}

	log.Println("[SEED]: Class group added: ", name)

}

func SeedSubject(db *gorm.DB) {
	name := "Wychowanie Fizyczne"
	subject := &domain.Subject{
		Name: name,
	}

	var wfTeacher domain.Teacher

	if err := db.FirstOrCreate(&wfTeacher, domain.Teacher{User: teacher}).Error; err != nil {
		log.Panic(err)
	}

	subject.Teachers = append(subject.Teachers, wfTeacher)

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
	log.Printf("[SEED]: Creating mark %f for subject %s", mark.Value, mark.Subject.Name)
}

func SeedSubjectTeacherClassGroup(db *gorm.DB) {
	var c domain.ClassGroup
	var t domain.Teacher
	var s domain.Subject

	if err := db.Find(&s, "id = ?", 1).Error; err != nil {
		panic(err)
	}

	if err := db.Find(&t, "id = ?", 1).Error; err != nil {
		panic(err)
	}

	if err := db.Find(&c, "id = ?", 1).Error; err != nil {
		panic(err)
	}

	if err := db.Create(&domain.SubjectTeacherClass{
		Teacher:    t,
		ClassGroup: c,
		Subject:    s,
	}).Error; err != nil {
		log.Panicln(err)
	}

	log.Println("[SEDD]: SeedSubjectTeacherClassGroup successfully")
}
