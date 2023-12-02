package db

import (
	"e-student/internal/app/domain"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"gorm.io/gorm"
)

func AskForCleanDatabase(db *gorm.DB) {
	if yes() {
		db.Migrator().DropTable(&domain.SubjectClassGroup{})
		db.Migrator().DropTable(&domain.SubjectTeacher{})
		db.Migrator().DropTable(&domain.Student{})
		db.Migrator().DropTable(&domain.Teacher{})
		db.Migrator().DropTable(&domain.User{})
		db.Migrator().DropTable(&domain.ClassGroup{})
		db.Migrator().DropTable(&domain.Mark{})
		db.Migrator().DropTable(&domain.Subject{})
	}
}

func yes() bool {
	prompt := promptui.Select{
		Label: "Would you like to cleanup database? All data will be lost [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result == "Yes"
}
