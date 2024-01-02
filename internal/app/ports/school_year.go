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

	SchoolYearRepository interface {
		GetAll() ([]domain.SchoolYear, error)
	}

	SchoolYearService interface {
		GetAll() ([]*SchoolYearDetailed, error)
	}
)
