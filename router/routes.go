package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func printMe() {
	println("init router")
}

func initializeRoutes(router *gin.Engine) {
	//Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix. For example, all the routes that use a common middleware for authorization could be grouped.
	v1 := router.Group("/api/v1")

	// Define a simple GET endpoint
	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.GET("/openings", func(ctx *gin.Context) {
		//statusCode, gin.H = transform in JSON
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET Openings",
		})
	})

	v1.POST("/opening", func(ctx *gin.Context) {
		//statusCode, gin.H = transform in JSON
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET Openings",
		})
	})

	v1.DELETE("/opening", func(ctx *gin.Context) {
		//statusCode, gin.H = transform in JSON
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET Openings",
		})
	})

	v1.PUT("/opening", func(ctx *gin.Context) {
		//statusCode, gin.H = transform in JSON
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET Openings",
		})
	})

	v1.GET("/opening", func(ctx *gin.Context) {
		//statusCode, gin.H = transform in JSON
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET Opening",
		})
	})

}
