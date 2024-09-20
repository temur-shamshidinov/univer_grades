package postgres

import (
	"context"
	"log"

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

	query := `
		INSERT INTO teachers(
			teacher_id,
			name,
			surname,
			email,
			password,
			created_at,
			updated_at
		) VALUES(
			$1, $2, $3, $4, $5, $6, $7 
		)
	`

	_, err := t.conn.Exec(
		ctx, query,
		req.TeacherID,
		req.Name,
		req.Surname,
		req.Email,
		req.Password,
		req.CreatedAt,
		req.UpdateAt,
	)

	if err != nil {
		log.Println("error with CreateTeacher", err)
		return err
	}

	return nil
}
func (t *teacherRepo) GetTeacherList(ctx context.Context, req *models.GetListReq) (*models.GetTeacherList, error) {

	limit := req.Limit
	page := req.Page

	page = (page - 1) * limit

	query := `
		SELECT 
			teacher_id,
			name,
			surname,
			email,
			created_at,
			updated_at
		FROM 
			teachers
		LIMIT
			$1
		OFFSET
			$2
	`

	rows, err := t.conn.Query(ctx, query, limit, page)
	if err != nil {
		log.Println("Error on GetTeacherList:", page)
		return nil, err
	}

	defer rows.Close()

	var teachers []*models.Teacher
	for rows.Next() {

		var teacher models.Teacher
		if err := rows.Scan(
			&teacher.TeacherID,
			&teacher.Name,
			&teacher.Surname,
			&teacher.Email,
			&teacher.CreatedAt,
			&teacher.UpdateAt,
		); err != nil {
			log.Println("Error on scaning teacher:", err)
			return nil, err
		}
		teachers = append(teachers, &teacher)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error on Rows:", err)
		return nil, err
	}

	return &models.GetTeacherList{
		Teachers: teachers,
		Count:    len(teachers),
	}, nil

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
