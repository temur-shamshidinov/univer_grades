package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/models"
	repoi "github.com/temur-shamshidinov/univer_grades/storage/repoI"
)

type groupRepo struct {
	conn *pgx.Conn
}

func NewGroupRepo(db *pgx.Conn) repoi.GroupRepoI {
	return &groupRepo{conn: db}
}

func (g *groupRepo) CreateGroup(ctx context.Context, req *models.Group) error {
	
	query := `
		INSERT INTO groups (
			group_id,
			group_name,
			course_id,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5 
		)
	`

	_, err := g.conn.Exec(
		ctx,query,
		req.GroupID,
		req.GroupName,
		req.CourseID,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		log.Println("Error with Create Group",err)
		return err
	}
	
	return nil
}

func (g *groupRepo) GetGroupList(ctx context.Context, req *models.GetListReq) (*models.GetGroupsList, error) {
	return nil, nil
}

func (g *groupRepo) GetGroupByID(ctx context.Context, id string) (*models.Group, error) {
	return nil, nil
}

func (g *groupRepo) UpdateGroup(ctx context.Context, req *models.Group) error {
	return nil
}

func (g *groupRepo) DeleteGroup(ctx context.Context, id string) error {
	return nil
}
