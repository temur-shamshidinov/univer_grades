package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/temur-shamshidinov/univer_grades/models"
	"github.com/temur-shamshidinov/univer_grades/pkg/helpers"
)

func (h handler) CreatStudent(ctx *gin.Context) {

	var reqBody models.Student
	var student = &models.Student{}

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		log.Println("Error on Bind reqBody, Create Student", err)
		ctx.JSON(400,err)
	}

	helpers.DataParser(reqBody,student)

	student.StudentID = uuid.New()
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	if err := h.storage.StudentRepo().CreateStudent(ctx,student); err != nil {
		log.Println("Error on CreateStudent",err)
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"created")
}