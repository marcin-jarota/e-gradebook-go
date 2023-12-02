package domain

import "gorm.io/gorm"

type SubjectClassGroup struct {
	gorm.Model
	SubjectID    uint
	ClassGroupID uint
}
