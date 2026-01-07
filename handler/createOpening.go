package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	//statusCode, gin.H = transform in JSON
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post Openings",
	})
}
