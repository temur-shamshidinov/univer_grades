package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/models"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type gradeRepo struct {
	conn *pgx.Conn
}

func NewGradeRepo(db *pgx.Conn) repoi.GradeRepoI {
	return &gradeRepo{conn: db}
}

func (g *gradeRepo) CreateGrade(ctx context.Context, req *models.Grade) error {
	return nil
}

func (g *gradeRepo) GetGradeList(ctx context.Context, req *models.GetListReq) (*models.GetGradesList, error) {
	return nil, nil
}

func (g *gradeRepo) GetGradeByID(ctx context.Context, id string) (*models.Grade, error) {
	return nil, nil
}

func (g *gradeRepo) UpdateGrade(ctx context.Context, req *models.Grade) error {
	return nil
}

func (g *gradeRepo) DeleteGrade(ctx context.Context, id string) error {
	return nil
}
