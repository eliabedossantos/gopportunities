package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	//statusCode, gin.H = transform in JSON
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Post Openings",
	// })

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	logger.Infof("request received: %v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

}
