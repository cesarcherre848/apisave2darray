package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/cesarcherre848/apisave2darray/src/controllers" // Importa desde el directorio local
)

// SetupRoutes configura las rutas de la API
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/data") // Prefijo común para las rutas
	{
		api.GET("/ping", controllers.Ping)
		api.POST("/upload", controllers.UploadData)
	}
}