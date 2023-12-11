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

		db.Migrator().DropTable(&domain.Student{})
		db.Migrator().DropTable(&domain.Teacher{})
		db.Migrator().DropTable(&domain.User{})
		db.Migrator().DropTable(&domain.ClassGroup{})
		db.Migrator().DropTable(&domain.Mark{})
		db.Migrator().DropTable(&domain.Subject{})
		db.Exec("DROP TABLE subject_teachers cascade")
		db.Exec("DROP TABLE subject_teacher_classes cascade")
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
