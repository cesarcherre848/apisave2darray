package main

import (
    "log"
    "github.com/gin-gonic/gin"
	"github.com/cesarcherre848/apisave2darray/src/routers"
    "github.com/cesarcherre848/apisave2darray/src/config"
    "github.com/cesarcherre848/apisave2darray/src/utils"

)

func main() {

	// Obtener el valor de la variable de entorno MODE o usar "debug" como valor por defecto
	mode := utils.GetEnvOrDefault("MODE", "debug")

	// Establecer el modo de Gin basado en el valor de la variable de entorno "MODE"
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode) // Establecer Gin en modo Release (producción)
	} else {
		gin.SetMode(gin.DebugMode) // Establecer Gin en modo Debug (desarrollo)
	}

	port := utils.GetEnvOrDefault("PORT", "2541") 
    config.Connect()
	defer func() {
		// Asegurarse de cerrar la conexión cuando la aplicación finalice
		config.Disconnect()
		log.Println("Conexión a la base de datos cerrada.")
	}()




    router := gin.Default()
	routers.SetupRoutes(router)


    // Ruta básica
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World !",
        })
    })

    // Inicia el servidor en el puerto 8080
    // Iniciar el servidor
	router.Run(":" + port)


}


