package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/models"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type studentRepo struct {
	conn *pgx.Conn
}

func NewStudentRepo(db *pgx.Conn) repoi.StudentRepoI {
	return &studentRepo{conn: db}
}

func (s *studentRepo) CreateStudent(ctx context.Context, req *models.Student) error {
	return nil
}

func (s *studentRepo) GetStudentList(ctx context.Context, req *models.GetListReq) (*models.GetStudentsList, error) {
	return nil, nil
}

func (s *studentRepo) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	return nil, nil
}

func (s *studentRepo) UpdateStudent(ctx context.Context, req *models.Student) error {
	return nil
}

func (s *studentRepo) DeleteStudent(ctx context.Context, id string) error {
	return nil
}
