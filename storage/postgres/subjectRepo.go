package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/models"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type subjectRepo struct {
	conn *pgx.Conn
}

func NewSubjectRepo(db *pgx.Conn) repoi.SubjectRepoI {
	return &subjectRepo{conn: db}
}

func (s *subjectRepo) CreateSubject(ctx context.Context, req *models.Subject) error {
	return nil
}

func (s *subjectRepo) GetSubjectList(ctx context.Context, req *models.GetListReq) (*models.GetSubjectsList, error) {
	return nil, nil
}

func (s *subjectRepo) GetSubjectByID(ctx context.Context, id string) (*models.Subject, error) {
	return nil, nil
}

func (s *subjectRepo) UpdateSubject(ctx context.Context, req *models.Subject) error {
	return nil
}

func (s *subjectRepo) DeleteSubject(ctx context.Context, id string) error {
	return nil
}
