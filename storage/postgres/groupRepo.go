package postgres

import (
	"context"

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
