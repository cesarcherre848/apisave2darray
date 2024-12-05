#!/bin/bash

# Definir las rutas
BIN_DIR="./bin"
BUILD_NAME="api-save2darray"

# Crear el directorio bin si no existe
if [ ! -d "$BIN_DIR" ]; then
    echo "Creando directorio bin..."
    mkdir "$BIN_DIR"
fi

# Compilar el código y mover el binario a la carpeta bin
echo "Compilando el proyecto..."
GOOS=linux GOARCH=amd64 go build -o "$BIN_DIR/$BUILD_NAME" ./src/main.go

if [ $? -eq 0 ]; then
    echo "Compilación exitosa. El binario está en $BIN_DIR/$BUILD_NAME"
else
    echo "Hubo un error en la compilación"
    exit 1
fi
