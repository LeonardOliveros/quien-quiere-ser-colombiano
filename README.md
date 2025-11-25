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

## 🛠️ Desarrollo

### Scripts Disponibles

#### Usando Make (Recomendado)
```bash
make help                  # Ver todos los comandos disponibles
make setup                 # Setup inicial completo (instala todo)
make dev-full             # Desarrollo: backend + frontend simultáneamente
make run                  # Solo backend
make frontend-dev         # Solo frontend en desarrollo
make build-all            # Build completo (frontend + backend)
make clean-sessions       # Limpiar todas las sesiones de juego
make clean-old-sessions   # Limpiar sesiones con datos antiguos
make db-stats             # Ver estadísticas de la base de datos
```

#### Usando NPM (desde root)
```bash
npm run dev               # Desarrollo: backend + frontend
npm run build             # Build completo
npm run frontend          # Solo frontend
npm run backend           # Solo backend
npm run clean:sessions    # Limpiar sesiones
npm run clean:db          # Eliminar base de datos
```

#### Frontend (desde /frontend)
```bash
npm run dev               # Servidor de desarrollo (Vite)
npm run build             # Build de producción
npm run build:watch       # Build con auto-reload
npm run type-check        # Verificar tipos TypeScript
npm run preview           # Preview del build
```

### Flujo de Desarrollo

1. **Setup inicial**
   ```bash
   make setup
   # o
   npm install && cd frontend && npm install
   ```

2. **Desarrollo con hot-reload**
   ```bash
   make dev-full
   # o
   npm run dev
   ```
   - Backend: http://localhost:8080
   - Frontend: http://localhost:5173

3. **Build para producción**
   ```bash
   make build-all
   # o
   npm run build
   ```

### Estructura del Proyecto
```
quiz/
├── frontend/              # Vue 3 + TypeScript + Vite
│   ├── src/
│   │   ├── views/        # Vistas: Login, Menu, Game, Results
│   │   ├── stores/       # Pinia stores (game, user)
│   │   ├── services/     # API services
│   │   └── types/        # TypeScript types
│   └── dist/             # Build de producción
├── *.go                  # Backend Go
├── handlers.go           # API handlers
├── models.go             # Modelos de datos
├── seeder.go             # Seed de preguntas
├── quiz.db               # SQLite database
├── Makefile              # Make commands
└── package.json          # NPM scripts
```

### Stack Tecnológico

**Backend:**
- Go 1.21+
- Gin (web framework)
- GORM (ORM)
- SQLite (base de datos)

**Frontend:**
- Vue 3 (Composition API)
- TypeScript
- Pinia (state management)
- Vue Router
- Axios
- Vite (build tool)
- Bootstrap 5

### Gestión de Base de Datos

```bash
# Ver estadísticas
make db-stats

# Limpiar sesiones antiguas (antes de actualización)
make clean-old-sessions

# Limpiar todas las sesiones
make clean-sessions

# Resetear base de datos completa
make reset-db
```

### Troubleshooting

**Problema: "Pregunta 1 de 753" en modo TIMED**
```bash
# Solución: Limpiar sesiones antiguas y crear nueva partida
make clean-old-sessions
# Luego en el navegador: Cmd+Shift+R (Mac) o Ctrl+Shift+R (Windows)
```

**Problema: Frontend no refleja cambios**
```bash
# Rebuild frontend
cd frontend && npm run build
# Limpiar cache del navegador: Cmd+Shift+R
```

**Problema: Base de datos corrupta**
```bash
make reset-db
```

## 🙏 Créditos

- Preguntas basadas en el documento oficial "COLOMBIA: NUESTRA CASA"
- Interfaz inspirada en "¿Quién quiere ser millonario?"
- Backend: Go + Gin + GORM + SQLite
- Frontend: Vue 3 + TypeScript + Vite + Pinia

---

**¡Mucha suerte en tu examen de naturalización! 🎓🇨🇴**
