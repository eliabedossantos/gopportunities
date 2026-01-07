package router

//sempre que você não especificar dessa forma: import (gin "github.com/gin-gonic/gin"), o nome da variavel do package, será utilizado o nome do ultimo modulo no caso o /gin
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// para exportar algo de um pacote, PRECISA iniciar comletras maiuscula senão ficará so no escopodo sub-package
func Initialize() {

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run(":8080")
}
