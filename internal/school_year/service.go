package schoolyear

import (
	"e-student/internal/app/ports"
)

type SchoolYearService struct {
	repo ports.SchoolYearRepository
}

func NewSchoolYearService(repo ports.SchoolYearRepository) *SchoolYearService {
	return &SchoolYearService{
		repo: repo,
	}
}

func (r *SchoolYearService) GetAll() ([]*ports.SchoolYearDetailed, error) {
	var schoolYears []*ports.SchoolYearDetailed
	years, err := r.repo.GetAll()

	if err != nil {
		return nil, err
	}
	for _, sr := range years {
		schoolYears = append(schoolYears, &ports.SchoolYearDetailed{
			SchoolYearBasic: ports.SchoolYearBasic{
				ID:   int(sr.ID),
				Name: sr.Name,
			},
			ClassGroupsCount: len(sr.ClassGroups),
			Start:            sr.Start,
			End:              sr.End,
			IsCurrent:        sr.IsCurrent,
		})
	}

	return schoolYears, nil
}
