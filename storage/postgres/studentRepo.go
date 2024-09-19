package postgres

import (
	"context"
	"log"

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

	query := `
		INSERT INTO students(
			student_id,
			name,
			surname,
			group_id,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6 
		)
	`

	_, err := s.conn.Exec(
		ctx,query,
		req.StudentID,
		req.Name,
		req.Surname,
		req.GroupID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	
	if err != nil {
		log.Println("Error with Create Student",err)
		return err
	}

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
