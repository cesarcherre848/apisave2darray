package config

import (
	"context"
	"fmt"
	"log"
	//"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"


	"github.com/cesarcherre848/apisave2darray/src/utils"
)

var Client *mongo.Client
var Database *mongo.Database

// Connect establece una conexión a la base de datos MongoDB
func Connect() {
	// Obtener valores desde variables de entorno
	username := utils.GetEnvOrDefault("MONGO_USERNAME", "root")
	password := utils.GetEnvOrDefault("MONGO_PASSWORD", "Mc05071995..")
	host := utils.GetEnvOrDefault("MONGO_HOST", "localhost")
	port := utils.GetEnvOrDefault("MONGO_PORT", "27017")
	databaseName := utils.GetEnvOrDefault("MONGO_DATABASE", "Condestable")  // Nombre de la base de datos predeterminada

	// Validar variables esenciales
	if username == "" || password == "" || host == "" || port == "" {
		log.Fatalf("Error: Las variables de entorno MONGO_USERNAME, MONGO_PASSWORD, MONGO_HOST y MONGO_PORT son obligatorias.")
	}

	// Construir la URI de conexión
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	// Establecer las opciones del cliente
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(10).       // Limita el tamaño del pool de conexiones
		SetRetryWrites(true).     // Activa reintentos automáticos para escrituras
		SetConnectTimeout(10 * time.Second) // Tiempo máximo de espera para conectar

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar conexión inicial
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	// Establecer la base de datos predeterminada
	if databaseName == "" {
		databaseName = "test" // Valor predeterminado si no se define
	}
	Database = Client.Database(databaseName)

	fmt.Println("Conexión a MongoDB exitosa")
}

// Reconnect intenta restablecer la conexión manualmente si se pierde
func Reconnect() {
	for {
		err := Client.Ping(context.Background(), nil)
		if err != nil {
			fmt.Println("Conexión perdida. Intentando reconectar...")
			Connect()
		} else {
			fmt.Println("Conexión a MongoDB activa")
		}
		time.Sleep(10 * time.Second) // Verifica la conexión cada 10 segundos
	}
}

// Disconnect cierra la conexión con la base de datos MongoDB
func Disconnect() {
	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error al desconectar de MongoDB: %v", err)
	}
	fmt.Println("Desconexión de MongoDB exitosa")
}
