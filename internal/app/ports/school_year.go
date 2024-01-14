package ports

import (
	"e-student/internal/app/domain"
	"time"
)

type (
	SchoolYearBasic struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	SchoolYearDetailed struct {
		SchoolYearBasic
		ClassGroupsCount int       `json:"classGroupsCount"`
		Start            time.Time `json:"start"`
		End              time.Time `json:"end"`
		IsCurrent        bool      `json:"isCurrent"`
	}

	SchoolYearPayload struct {
		Name  string `json:"name"`
		Start string `json:"start"`
		End   string `json:"end"`
	}

	SchoolYearRepository interface {
		GetAll() ([]domain.SchoolYear, error)
		AddSchoolYear(name string, start time.Time, end time.Time) error
	}

	SchoolYearService interface {
		GetAll() ([]*SchoolYearDetailed, error)
		AddSchoolYear(schoolYear SchoolYearPayload) error
	}
)
