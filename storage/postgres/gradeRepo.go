package postgres

import (
	"context"
	"log"

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
	
	query := `
		INSERT INTO grades(
			grade_id,
			grade_name,
			grade_value,
			grade_date,
			subject_id,
			group_id,
			student_id,
			created_id,
			updated_id
		) VALUES(
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`
	
	_, err := g.conn.Exec(
		ctx,query,
		req.GradeID,
		req.GradeName,
		req.GradeValue,
		req.GradeDate,
		req.SubjectID,
		req.GroupID,
		req.StudentID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("Error with Create Grade", err)
		return err
	}
	
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
