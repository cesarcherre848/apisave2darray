package models

import (
	"context"
	"time"

	"github.com/cesarcherre848/apisave2darray/src/config"
	"github.com/google/uuid"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	ID          string    `json:"id" bson:"_id"`         // ID único
	Measurement string    `json:"measurement" bson:"measurement"`
	Timestamp   float64   `json:"timestamp" bson:"timestamp"`
	X           []float64     `json:"x" bson:"x"`
	Y           []float64     `json:"y" bson:"y"`
	X_Unit      string    `json:"x_unit" bson:"x_unit"`
	Y_Unit      string    `json:"y_unit" bson:"y_unit"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

func InsertData(data Data) (*mongo.InsertOneResult, error) {
	collection := config.Database.Collection("DataMeasurement")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Generar un UUID para el campo ID
	data.ID = uuid.New().String()
	data.CreatedAt = time.Now()

	// Convertir la estructura a BSON
	result, err := collection.InsertOne(ctx, data)
	return result, err
}