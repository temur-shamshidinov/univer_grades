package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/temur-shamshidinov/univer_grades/models"
	"github.com/temur-shamshidinov/univer_grades/pkg/helpers"
)

func (h handler) CreatGrade(ctx *gin.Context) {

	var reqBody models.Grade
	var grade = &models.Grade{}

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		log.Println("Error on Bind reqBody, Creare Grade", err)
		ctx.JSON(400, err)
	}

	helpers.DataParser(reqBody, grade)

	grade.GradeID = uuid.New()
	grade.GradeDate = time.Now()
	grade.CreatedAt = time.Now()
	grade.UpdatedAt = time.Now()

	if err := h.storage.GradeRepo().CreateGrade(ctx,grade); err != nil {
		log.Println("Error on CreateGrade",err)
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"created")
}
