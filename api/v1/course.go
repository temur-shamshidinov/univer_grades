package v1

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/temur-shamshidinov/univer_grades/models"
	"github.com/temur-shamshidinov/univer_grades/pkg/helpers"
)



func (h handler) CreateCourse(ctx *gin.Context) {

	var reqBody models.Course
	var course = &models.Course{}

	err := ctx.Bind(&reqBody)
	if err != nil {
		log.Println("Error Bind reqBody CreateCourse:", err)
		ctx.JSON(400, err)
	}

	helpers.DataParser(reqBody, course)

	course.CourseID = uuid.New()
	course.CreatedAt = time.Now()
	course.UpdateAt = time.Now()

	if err := h.storage.CourseRepo().CreateCourse(ctx,course); err != nil {
			log.Println("Error on CreatCourse:", err) 
			ctx.JSON(500, err)
		return
	}
	
	ctx.JSON(201, gin.H{"message": "course created successfully", "course_id": course.CourseID})

}
