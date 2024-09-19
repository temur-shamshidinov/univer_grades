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

	router.Run(":8080")

}
