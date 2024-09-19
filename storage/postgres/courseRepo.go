package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/models"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type courseRepo struct {
	conn *pgx.Conn
}

func NewCourseRepo(db *pgx.Conn) repoi.CourseRepoI {
	return &courseRepo{conn: db}
}

func (c *courseRepo) CreateCourse(ctx context.Context, req *models.Course) error {

	query := `
		INSERT INTO courses(
			course_id,
			course_name,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4 
		)
	`
	_, err := c.conn.Exec(
		ctx, query,
		req.CourseID,
		req.CourseName,
		req.CreatedAt,
		req.UpdateAt,
	)
	if err != nil {
		log.Println("Error with Create Course", err)
		return err
	}

	return nil
}

func (c *courseRepo) GetCourseList(ctx context.Context, req *models.GetListReq) (*models.GetCoursesList, error) {
	return nil, nil
}

func (c *courseRepo) GetCourseByID(ctx context.Context, id string) (*models.Course, error) {
	return nil, nil
}

func (c *courseRepo) UpdateCourse(ctx context.Context, req *models.Course) error {
	return nil
}

func (c *courseRepo) DeleteCourse(ctx context.Context, id string) error {
	return nil
}
