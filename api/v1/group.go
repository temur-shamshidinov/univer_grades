package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/temur-shamshidinov/univer_grades/models"
	"github.com/temur-shamshidinov/univer_grades/pkg/helpers"
)

func (h handler) CreatGroup(ctx *gin.Context) {

	var reqBody models.Group
	var group = &models.Group{}

	err := ctx.Bind(&reqBody)
	if err != nil {
		log.Println("Error Bind req Body, Create Group", err)
		ctx.JSON(400, err)
	}

	helpers.DataParser(reqBody, group)

	group.GroupID = uuid.New()
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	if err := h.storage.GroupRepo().CreateGroup(ctx, group); err != nil {
		log.Println("Error on CreateGroup", err)
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, "created")

}
