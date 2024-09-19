package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/temur-shamshidinov/univer_grades/models"
	"github.com/temur-shamshidinov/univer_grades/pkg/helpers"
)

func (h handler) CreateSubject(ctx *gin.Context) {

	var reqBody models.Subject
	var subject = &models.Subject{}

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		log.Println("Error Bind reqBody, Create Subject", err)
		ctx.JSON(400,err)
	}

	helpers.DataParser(reqBody,subject)

	subject.SubjectID = uuid.New()
	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()

	if err := h.storage.SubjectRepo().CreateSubject(ctx,subject); err != nil {
		log.Println("Error on CreateGroup", err)
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"created")
}