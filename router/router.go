package router

//sempre que você não especificar dessa forma: import (gin "github.com/gin-gonic/gin"), o nome da variavel do package, será utilizado o nome do ultimo modulo no caso o /gin
import (
	"github.com/gin-gonic/gin"
)

// para exportar algo de um pacote, PRECISA iniciar comletras maiuscula senão ficará so no escopodo sub-package
func Initialize() {
	// voce pode acessar funcoes/variaveis de outros arquivos do mesmo package sem precisar importa-las
	printMe()

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	//initilize Routes
	initializeRoutes(router)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run(":8080")
}
