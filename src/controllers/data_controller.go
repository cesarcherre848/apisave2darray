package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cesarcherre848/apisave2darray/src/models"
)

// Ping es una ruta básica de prueba
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}


func UploadData(c *gin.Context) {
	// Variable para capturar el JSON recibido
	var dataArray []models.Data // Usamos el tipo Data del modelo

	// Intentar enlazar el cuerpo de la solicitud al array de Data
	if err := c.ShouldBindJSON(&dataArray); err != nil {
		// Si hay un error en el enlace, devolver un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que los arreglos "x" y "y" tengan la misma longitud
	for _, data := range dataArray {
		if len(data.X) != len(data.Y) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Arrays 'x' and 'y' must have the same length"})
			return
		}
		// Generar un ID único y la fecha de creación para cada entrada
		//data.ID = uuid.New().String()
		//data.CreatedAt = time.Now()
	}

	// Realizar la inserción en la base de datos asincrónicamente
	go func() {
		for _, data := range dataArray {
			// Insertar cada dato en la base de datos
			_, err := models.InsertData(data)
			if err != nil {
				// Si hay error en la inserción, loguearlo
				// Puedes agregar un logger o simplemente registrar el error
				// log.Printf("Error al insertar dato: %v\n", err)
			}
		}
	}()

	// Responder al cliente de inmediato después de que todo se haya procesado
	c.JSON(http.StatusOK, gin.H{
		"message": "Data received successfully, processing in background",
	})
}