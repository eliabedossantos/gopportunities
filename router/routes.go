package router

import (
	"github.com/eliabedossantos/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func printMe() {
	println("init router")
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func initializeRoutes(router *gin.Engine) {
	//Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix. For example, all the routes that use a common middleware for authorization could be grouped.
	v1 := router.Group("/api/v1")

	v1.GET("/openings", handler.ListOpeningsHandler)
	v1.POST("/opening", handler.ShowOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.GET("/opening", handler.ShowOpeningHandler)

	// v1.GET("/n8nTest", func(ctx *gin.Context) {
	// 	products := []Product{
	// 		{ID: 1, Name: "maquiagem", Value: 28.23},
	// 		{ID: 2, Name: "creme de pentear", Value: 12.34},
	// 	}

	// 	ctx.JSON(200, gin.H{
	// 		"data": products,
	// 	})
	// })
}
