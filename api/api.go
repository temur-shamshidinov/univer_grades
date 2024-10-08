package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/temur-shamshidinov/univer_grades/api/v1"
	"github.com/temur-shamshidinov/univer_grades/storage"
)

func Api(storage storage.StorageI) {

	router := gin.Default()

	h := v1.NewHandler(storage)

	router.GET("/ping", h.Ping)

	// teacher

	router.POST("/create-teacher", h.CreateTeacher)

	// course

	router.POST("/create-course", h.CreateCourse)

	// group

	router.POST("/create-group",h.CreatGroup)

	// subject

	router.POST("/create-subject", h.CreateSubject)

	// student

	router.POST("/create-student",h.CreatStudent)

	// grade

	router.POST("/create-grade",h.CreatGrade)

	router.Run(":8080")

}
