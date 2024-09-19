package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/storage/postgres"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type StorageI interface {
	TeacherRepo() repoi.TeacherRepoI
}

type storage struct {
	teacherRepo repoi.TeacherRepoI
}

func NewStorage(db *pgx.Conn) StorageI {
	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
	}
}

func (s storage) TeacherRepo() repoi.TeacherRepoI {
	return s.teacherRepo
}
