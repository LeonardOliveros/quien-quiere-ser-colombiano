# Quiz App - ¿Quién Quiere Ser Colombiano? 🇨🇴

Una aplicación interactiva estilo "Quién quiere ser millonario" para preparar el examen de naturalización colombiana.

## 🎮 Características

### Modos de Juego
- **Modo Práctica**: Sin límite de tiempo, aprende a tu ritmo
- **Contrarreloj**: 80 preguntas en 3 horas, simula el examen real
- **Áreas Débiles**: Enfoque inteligente en tus áreas de mejora
- **Por Categoría**: Practica temas específicos

### Categorías
- 🏛️ **Constitución Política** (60% mínimo para aprobar)
- 🌎 **Geografía** (55% mínimo para aprobar)
- 📚 **Historia Patria** (40% mínimo para aprobar)
- 🎭 **Cultura y Sociedad** (40% mínimo para aprobar)

### Funcionalidades
- ✅ Base de datos con 450+ preguntas del examen oficial
- 📊 Estadísticas detalladas por categoría
- 🚩 Sistema de marcado de preguntas dudosas
- 💡 Ayudas estilo millonario (50:50, Pista, Saltar)
- 📈 Seguimiento de progreso histórico
- 🎯 Recomendaciones personalizadas de estudio
- 💾 Almacenamiento de todas las partidas
- 🏆 Sistema de puntuación y logros

## 🚀 Instalación

### Requisitos
- Go 1.21 o superior
- Git

### Pasos de instalación

1. **Clonar o copiar el proyecto**
```bash
mkdir quiz-app
cd quiz-app
# Copiar todos los archivos del proyecto aquí
```

2. **Instalar dependencias**
```bash
go mod download
```

3. **Ejecutar la aplicación**
```bash
go run .
```

4. **Abrir en el navegador**
```
http://localhost:8080
```

## 📁 Estructura del Proyecto

```
quiz-app/
├── main.go              # Archivo principal
├── models.go            # Modelos de datos
├── handlers.go          # Controladores API
├── seeder.go            # Poblador de preguntas
├── go.mod               # Dependencias Go
├── quiz.db              # Base de datos SQLite (se crea automáticamente)
├── templates/
│   └── index.html       # Interfaz web principal
└── static/
    ├── style.css        # Estilos CSS
    └── app.js           # Lógica JavaScript
```

## 💻 Uso de la Aplicación

### Primer Uso
1. Al iniciar la aplicación por primera vez, se creará la base de datos
2. Se poblarán automáticamente las preguntas iniciales
3. Registra tu usuario o inicia sesión

### Durante el Juego
- **Responder**: Click en la opción que consideres correcta
- **Marcar pregunta**: Usa el botón "Marcar" si no estás seguro
- **Ayudas**: Usa las ayudas disponibles estratégicamente
- **Pausar**: Puedes pausar en cualquier momento (excepto en modo contrarreloj)

### Después del Juego
- Revisa tus respuestas incorrectas con explicaciones
- Consulta las preguntas marcadas
- Lee las recomendaciones de estudio
- Verifica tu progreso por categoría

## 🎯 Estrategias de Estudio

1. **Comienza con Modo Práctica** para familiarizarte con las preguntas
2. **Usa el modo Por Categoría** para fortalecer áreas específicas
3. **Activa Áreas Débiles** cuando tengas identificadas tus falencias
4. **Practica con Contrarreloj** cuando te sientas preparado

## 📊 Sistema de Puntuación

- ✅ Respuesta correcta: +10 puntos
- ❌ Respuesta incorrecta: 0 puntos
- ⏱️ Bonus por velocidad en modo contrarreloj
- 🎯 Multiplicadores por rachas correctas

## 🔧 Configuración Avanzada

### Variables de Entorno (.env)
```env
PORT=8080                    # Puerto del servidor
DATABASE_PATH=quiz.db        # Ruta de la base de datos
```

### Agregar Más Preguntas

Edita el archivo `seeder.go` y agrega preguntas en el formato:
```go
{
    Category:    "CATEGORIA",
    SubCategory: "Subcategoría",
    Text:        "¿Pregunta?",
    Difficulty:  1-5,
    Points:      10,
    Explanation: "Explicación de la respuesta",
    Choices: []Choice{
        {Text: "Opción A", IsCorrect: false, Order: 1},
        {Text: "Opción B", IsCorrect: true, Order: 2},
        {Text: "Opción C", IsCorrect: false, Order: 3},
        {Text: "Opción D", IsCorrect: false, Order: 4},
    },
}
```

## 🐛 Solución de Problemas

### La aplicación no inicia
- Verifica que el puerto 8080 esté disponible
- Asegúrate de tener permisos de escritura para crear la base de datos

### Las preguntas no cargan
- Elimina `quiz.db` y reinicia la aplicación
- Verifica que `seeder.go` tenga las preguntas correctamente formateadas

### Error de dependencias
```bash
go mod tidy
go mod download
```

## 📝 Notas Importantes

- La base de datos se crea automáticamente en el primer inicio
- Las sesiones se mantienen en el navegador (localStorage)
- El progreso se guarda automáticamente
- Requiere conexión a internet para cargar librerías CSS/JS externas

## 🤝 Contribuciones

Para agregar más preguntas o mejorar la aplicación:
1. Edita los archivos correspondientes
2. Prueba los cambios localmente
3. Documenta las mejoras

## 📄 Licencia

Proyecto educativo para preparación del examen de naturalización colombiana.

## 🙏 Créditos

- Preguntas basadas en el documento oficial "COLOMBIA: NUESTRA CASA"
- Interfaz inspirada en "¿Quién quiere ser millonario?"
- Desarrollado con Go, Gin, GORM, y tecnologías web modernas

---

**¡Mucha suerte en tu examen de naturalización! 🎓🇨🇴**
