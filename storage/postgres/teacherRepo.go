package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/models"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type teacherRepo struct {
	conn *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn) repoi.TeacherRepoI {
	return &teacherRepo{conn: db}
}

func (t *teacherRepo) CreateTeacher(ctx context.Context, req *models.Teacher) error {

	
	return nil
}
func (t *teacherRepo) GetTeacherList(ctx context.Context, req *models.GetListReq) (*models.GetTeacherList, error) {
	return nil, nil
}
func (t *teacherRepo) GetTeacherByID(ctx context.Context, id string) (*models.Teacher, error) {
	return nil, nil
}
func (t *teacherRepo) UpadeTeacher(ctx context.Context, req *models.Teacher) error {
	return nil
}
func (t *teacherRepo) DeletedTeacher(ctx context.Context, id string) error {
	return nil
}
