package db

import (
	"e-student/internal/app/domain"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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

func SeedSchoolYear(db *gorm.DB) {
	if err := db.Exec(`SELECT open_new_school_year('2023/2024', '2023-09-01', '2024-06-30')`).Error; err != nil {
		log.Fatal(err)
	}

	log.Println("[SEED]: Creating domain.SchoolYear added: ", "2023/2024")
}

func SeedStudentUser(db *gorm.DB) {
	hash, err := student.GeneratePassword(student.Password)

	if err != nil {
		panic(err)
	}

	var group domain.ClassGroup

	err = db.FirstOrCreate(&group, domain.ClassGroup{Name: "1 mat-fiz"}).Error

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

	if err := db.FirstOrCreate(&classGroup, domain.ClassGroup{Name: name, EducationYear: 1}).Error; err != nil {
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
			log.Println("[SEED]: Creating domain.Subject: ", name)
		} else {
			panic(err)
		}
	}
}

func SeedLessons(db *gorm.DB) {
	name := "1 mat-fiz"
	var classGroup domain.ClassGroup

	if err := db.FirstOrCreate(&classGroup, domain.ClassGroup{Name: name}).Error; err != nil {
		log.Fatal(err)
	}

	lesson := domain.Lesson{
		TeacherID:    1,
		SubjectID:    1,
		ClassGroupID: classGroup.ID,
		Start:        time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.UTC),
		End:          time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 45, 0, 0, time.UTC),
		DayOfWeek:    1,
	}

	if err := db.Create(&lesson).Error; err != nil {
		log.Fatal(err)
	}

	log.Println("[SEED]: Created domain.Lesson: ", lesson)
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

	if err := db.Create(&mark).Error; err != nil {
		log.Panic(err)
	}

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

	log.Println("[SEED]: SeedSubjectTeacherClassGroup successfully")
}

func LoadSQLFile(fileName string, db *gorm.DB) {
	query, err := loadSQL(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Exec(query).Error; err != nil {
		log.Fatal(err)
	}

	log.Printf("[SEED]: Loaded %s", fileName)
}

func loadSQL(fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%s/db/%s", cwd, fileName)

	sqlFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	sqlBytes, err := ioutil.ReadAll(sqlFile)
	if err != nil {
		return "", err
	}
	defer sqlFile.Close()

	return string(sqlBytes), nil
}
