package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/storage/postgres"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type StorageI interface {
	TeacherRepo() repoi.TeacherRepoI
	CourseRepo() repoi.CourseRepoI
	GroupRepo() repoi.GroupRepoI
	SubjectRepo() repoi.SubjectRepoI
	StudentRepo() repoi.StudentRepoI
	GradeRepo() repoi.GradeRepoI
}

type storage struct {
	teacherRepo repoi.TeacherRepoI
	courseRepo  repoi.CourseRepoI
	groupRepo   repoi.GroupRepoI
	subjectRepo repoi.SubjectRepoI
	studentRepo repoi.StudentRepoI
	gradeRepo   repoi.GradeRepoI
}

func NewStorage(db *pgx.Conn) StorageI {
	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
		courseRepo:  postgres.NewCourseRepo(db),
		groupRepo: postgres.NewGroupRepo(db),
		subjectRepo: postgres.NewSubjectRepo(db),
		studentRepo: postgres.NewStudentRepo(db),
		gradeRepo: postgres.NewGradeRepo(db),
	}
}

func (s storage) TeacherRepo() repoi.TeacherRepoI {
	return s.teacherRepo
}

func (s storage) CourseRepo() repoi.CourseRepoI{
	return s.courseRepo
}

func (s storage) GroupRepo() repoi.GroupRepoI {
	return s.groupRepo
}

func (s storage) SubjectRepo() repoi.SubjectRepoI{
	return s.subjectRepo
}
func (s storage) StudentRepo() repoi.StudentRepoI {
	return s.studentRepo
}

func (s storage) GradeRepo() repoi.GradeRepoI{
	return s.gradeRepo
}