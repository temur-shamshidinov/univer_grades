package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/temur-shamshidinov/univer_grades/models"
	"github.com/temur-shamshidinov/univer_grades/pkg/helpers"
)

func (h handler) CreateTeacher(ctx *gin.Context) {

	var reqBody models.TeacherCreateReq
	var teacher = &models.Teacher{}

	err := ctx.BindJSON(&reqBody)

	if err != nil {
		log.Println("Error Bind reqBody CreateTeacher:", err)
		ctx.JSON(400, err)
	}

	helpers.DataParser(reqBody, teacher)

	teacher.TeacherID = uuid.New()
	teacher.CreatedAt = time.Now()
	teacher.UpdateAt = time.Now()

	if err := h.storage.TeacherRepo().CreateTeacher(ctx,teacher); err != nil {
		log.Println("Error on CreateTeacher:", err)
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"created")

}
