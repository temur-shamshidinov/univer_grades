package repoi

import (
	"context"

	"github.com/temur-shamshidinov/univer_grades/models"
)

type TeacherRepoI interface{
	CreateTeacher(ctx context.Context, req *models.Teacher) error
	GetTeacherList(ctx context.Context, req *models.GetListReq) (*models.GetTeacherList, error)
	GetTeacherByID(ctx context.Context, id string) (*models.Teacher, error)
	UpadeTeacher(ctx context.Context, req *models.Teacher) error
	DeletedTeacher(ctx context.Context, id string) error
}