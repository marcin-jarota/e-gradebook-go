package domain

type SubjectTeacherClass struct {
	Teacher      Teacher
	TeacherID    int
	Subject      Subject
	SubjectID    int
	ClassGroup   ClassGroup
	ClassGroupID int
}
