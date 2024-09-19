package postgres

import (
	"context"
	"log"

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

	query := `
	
		INSERT INTO subjects (
			subject_id,
			subject_name,
			group_id,
			teacher_id,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6 
		)
	`

	_, err := s.conn.Exec(
		ctx,query,
		req.SubjectID,
		req.SubjectName,
		req.GroupID,
		req.TeacherID,
		req.CreatedAt,
		req.UpdatedAt,
	)

	if err != nil {
		log.Println("error with Create Subject", err)
		return err
	}

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
