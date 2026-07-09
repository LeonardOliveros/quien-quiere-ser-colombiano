#!/bin/bash

# Quiz App - Start Script

echo "======================================"
echo "Quiz de Naturalización Colombia"
echo "¿Quién Quiere Ser Colombiano?"
echo "======================================"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go no está instalado. Por favor instala Go 1.21 o superior."
    echo "Visita: https://golang.org/dl/"
    exit 1
fi

echo "✅ Go instalado: $(go version)"

# Download dependencies
echo ""
echo "📦 Descargando dependencias..."
go mod download

if [ $? -ne 0 ]; then
    echo "❌ Error al descargar dependencias"
    exit 1
fi

echo "✅ Dependencias instaladas"

# Check if database exists
if [ -f "quiz.db" ]; then
    echo "✅ Base de datos encontrada"
else
    echo "📝 Se creará una nueva base de datos al iniciar"
fi

# Start the application
echo ""
echo "🚀 Iniciando aplicación..."
echo "📍 La aplicación estará disponible en: http://localhost:8080"
echo ""
echo "Presiona Ctrl+C para detener el servidor"
echo "======================================"
echo ""

# Run the application
go run .
