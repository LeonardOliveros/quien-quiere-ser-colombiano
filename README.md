# Quiz App - ¿Quién Quiere Ser Colombiano? 🇨🇴

Una aplicación interactiva estilo "Quién quiere ser millonario" para preparar el examen de naturalización colombiana.

## 🎮 Características

### Modos de Juego
- **Modo Práctica**: Sin límite de tiempo, aprende a tu ritmo
- **Contrarreloj**: 80 preguntas (20 por categoría) en 1 hora, simula el examen real
- **Áreas Débiles**: Enfoque inteligente en tus áreas de mejora
- **Por Categoría**: Practica temas específicos

### Categorías
- 🏛️ **Constitución Política** (60% mínimo para aprobar)
- 🌎 **Geografía** (55% mínimo para aprobar)
- 📚 **Historia Patria** (40% mínimo para aprobar)
- 🎭 **Cultura y Sociedad** (40% mínimo para aprobar)

### Funcionalidades
- ✅ Base de datos con 750+ preguntas del examen oficial
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
├── main.go                        # Wiring: elige el adaptador de storage (DB_DRIVER) y arranca el server
├── handlers.go                    # Adaptador HTTP (Gin): solo habla con los puertos del dominio
├── internal/
│   ├── domain/                    # Núcleo hexagonal (sin dependencias de storage)
│   │   ├── models.go              # Entidades: Question, GameSession, User, ...
│   │   └── ports.go               # Puertos: Store, UserRepository, QuestionRepository, ...
│   ├── seed/                      # Carga y validación del banco de preguntas embebido
│   └── storage/
│       ├── sqlite/                # Adaptador SQLite/GORM (default, local)
│       ├── dynamodb/              # Adaptador DynamoDB (nube/serverless, single-table)
│       └── storagetest/           # Suite de conformidad compartida entre adaptadores
├── data/
│   ├── taxonomy.json              # Categorías y subcategorías canónicas
│   └── questions/                 # Banco de preguntas por categoría (embebido en el binario)
│       ├── cultura.json
│       ├── geografia.json
│       ├── historia.json
│       └── constitucion.json
├── go.mod                         # Dependencias Go
├── quiz.db                        # Base de datos SQLite (se crea automáticamente)
└── frontend/                      # SPA Vue 3 + TypeScript (build en dist/)
```

### Arquitectura hexagonal (puertos y adaptadores)

La persistencia está detrás del puerto `domain.Store` (`internal/domain/ports.go`):
los handlers HTTP nunca tocan SQL ni GORM, solo interfaces del dominio. El adaptador
se elige al arrancar con la variable `DB_DRIVER`:

- `DB_DRIVER=sqlite` (o vacío): `internal/storage/sqlite`, para desarrollo local.
- `DB_DRIVER=dynamodb`: `internal/storage/dynamodb`, para la nube/serverless.
  Usa una sola tabla (PK/SK, sin GSIs) con lecturas fuertemente consistentes;
  el diseño de claves está documentado en `internal/storage/dynamodb/keys.go`.

Ambos adaptadores pasan la misma suite de conformidad
(`internal/storage/storagetest`), que codifica la semántica exacta del puerto:
`go test ./internal/storage/...` la corre contra SQLite siempre, y contra
DynamoDB Local cuando `DYNAMODB_TEST_ENDPOINT` está definido
(`make dynamodb-local && make test-integration`).

Para agregar otro motor (Postgres, Turso, ...) basta con implementar `domain.Store`
en un paquete nuevo y registrarlo en `openStore()` de `main.go`.

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

Copia `.env.example` a `.env` y ajusta según necesites:
```env
PORT=8080                    # Puerto del servidor
DATABASE_PATH=quiz.db        # Ruta de la base de datos
GIN_MODE=debug               # debug | release | test
ALLOWED_ORIGINS=             # Orígenes CORS permitidos (vacío = todos, solo dev)
```

### Agregar Más Preguntas

Edita el archivo de la categoría correspondiente en `data/questions/` (embebido en el binario al compilar) y agrega preguntas en el formato:
```json
{
    "key": "CUL-0241",
    "subcategory": "GASTRONOMIA",
    "text": "¿Pregunta?",
    "difficulty": 2,
    "points": 10,
    "hint": "",
    "explanation": "Explicación de la respuesta",
    "choices": [
        {"text": "Opción A", "is_correct": false, "order": 1},
        {"text": "Opción B", "is_correct": true, "order": 2},
        {"text": "Opción C", "is_correct": false, "order": 3},
        {"text": "Opción D", "is_correct": false, "order": 4}
    ]
}
```

Reglas:
- `key` es el identificador estable de la pregunta (prefijo de categoría + consecutivo). No reutilices ni cambies keys existentes.
- `subcategory` debe ser un código definido en `data/taxonomy.json` para esa categoría. Para crear una subcategoría nueva, agrégala primero a la taxonomía.
- Cada pregunta debe tener 2+ opciones y exactamente una correcta; el servidor valida esto al arrancar y falla si no se cumple.

El seeder corre en cada arranque y sincroniza los archivos con la base de datos: crea preguntas nuevas y actualiza las modificadas (por `key`), sin duplicar. No hace falta borrar `quiz.db` para aplicar cambios.

## ☁️ Despliegue serverless en AWS

La app corre 100% serverless: Lambda (Go, arm64) + DynamoDB (on-demand) +
API Gateway HTTP API + S3 + CloudFront, definido con CDK en `infra/`.

```
Navegador ── CloudFront (HTTPS, un solo dominio)
               ├─ por defecto ──> S3 privado (OAC) ── SPA Vue (dist/)
               └─ /api/* ───────> API Gateway ──> Lambda ──> DynamoDB
```

Como el frontend usa rutas relativas (`/api`), no hay CORS ni cambios de
configuración: CloudFront enruta el mismo dominio a S3 y al API.

### Requisitos

- Cuenta AWS con credenciales configuradas (`aws configure`)
- Node.js 18+ (CDK) y Go 1.25+
- Primera vez en la cuenta/región: `cd infra && npx cdk bootstrap`

### Comandos

```bash
make infra-install   # instala dependencias de CDK (una vez)
make synth           # genera el template (compila la Lambda)
make deploy          # build del frontend + despliegue completo
make seed-remote     # re-sincroniza el banco de preguntas en la Lambda
make destroy         # elimina el stack (la tabla DynamoDB se conserva)
```

`make deploy` imprime el `DistributionUrl` (https://xxxx.cloudfront.net): esa
es la app. El banco de preguntas se siembra automáticamente al desplegar (un
custom resource invoca la Lambda con `{"quizapp_action":"seed"}`); en la
Lambda `SEED_ON_START=false`, así los cold starts no tocan el banco.

### Desarrollo local contra DynamoDB

```bash
make dynamodb-local    # DynamoDB Local en :8000 (Docker o java + DDB_LOCAL_DIR)
make test-integration  # suite de conformidad contra DynamoDB Local
make run-ddb           # la app local con DB_DRIVER=dynamodb (tabla auto-creada)
make seed-local        # siembra one-off (go run . -seed)
```

### CI/CD (GitHub Actions)

`.github/workflows/deploy.yml` corre `make lambda-build` + `cdk deploy` en
cada push a `main` (y manualmente vía "Run workflow"). Se autentica contra AWS
por OIDC — sin llaves de acceso de larga duración guardadas en GitHub.

Configuración única, con credenciales admin locales (`aws configure`):

```bash
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
REGION=us-east-1

# 1. Bootstrap de la cuenta/región (si no se hizo ya) — crea los roles
#    cdk-hnb659fds-*-role-$ACCOUNT_ID-$REGION que la CI va a asumir.
cd infra && npx cdk bootstrap "aws://$ACCOUNT_ID/$REGION" && cd ..

# 2. Proveedor OIDC de GitHub (una vez por cuenta AWS; puede que ya exista)
aws iam create-open-id-connect-provider \
  --url https://token.actions.githubusercontent.com \
  --client-id-list sts.amazonaws.com \
  --thumbprint-list 6938fd4d98bab03faadb97b34396831e3780aea

# 3. Rol que GitHub Actions asume, restringido a este repo/rama
cat > /tmp/trust-policy.json <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "Federated": "arn:aws:iam::$ACCOUNT_ID:oidc-provider/token.actions.githubusercontent.com" },
    "Action": "sts:AssumeRoleWithWebIdentity",
    "Condition": {
      "StringEquals": { "token.actions.githubusercontent.com:aud": "sts.amazonaws.com" },
      "StringLike": { "token.actions.githubusercontent.com:sub": "repo:LeonardOliveros/quien-quiere-ser-colombiano:ref:refs/heads/main" }
    }
  }]
}
EOF
aws iam create-role --role-name github-quiz-app-deploy \
  --assume-role-policy-document file:///tmp/trust-policy.json

# 4. Permisos: el rol solo necesita poder asumir los roles que cdk bootstrap
#    ya creó (deploy/file-publishing/lookup) — esos ya tienen los permisos
#    reales para tocar CloudFormation, Lambda, S3, DynamoDB, etc.
cat > /tmp/deploy-policy.json <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Action": "sts:AssumeRole",
    "Resource": "arn:aws:iam::$ACCOUNT_ID:role/cdk-hnb659fds-*-role-$ACCOUNT_ID-$REGION"
  }]
}
EOF
aws iam put-role-policy --role-name github-quiz-app-deploy \
  --policy-name cdk-assume-bootstrap-roles \
  --policy-document file:///tmp/deploy-policy.json

aws iam get-role --role-name github-quiz-app-deploy --query Role.Arn --output text
```

Guarda el ARN que imprime el último comando como secret del repo
`AWS_DEPLOY_ROLE_ARN` (Settings → Secrets and variables → Actions → Secrets).

### Costos y dominio propio

Todo es on-demand/free-tier friendly: sin tráfico no hay costo fijo salvo
centavos de S3/CloudFront. Para usar dominio propio (Route 53 + ACM), el stack
ya deja los puntos de enganche documentados en
`infra/lib/quiz-app-stack.ts` (`domainName`/`hostedZoneDomain`): certificado
DNS-validado, alias en CloudFront y un `ARecord` — sin reestructurar nada.

## 🐛 Solución de Problemas

### La aplicación no inicia
- Verifica que el puerto 8080 esté disponible
- Asegúrate de tener permisos de escritura para crear la base de datos

### Las preguntas no cargan
- Revisa el log de arranque: el seeder valida `data/taxonomy.json` y `data/questions/*.json` y reporta el error exacto (key duplicada, subcategoría inexistente, opciones inválidas)
- Como último recurso, elimina `quiz.db` y reinicia la aplicación

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
├── main.go               # Wiring y arranque
├── handlers.go           # Adaptador HTTP (Gin)
├── internal/             # Dominio, seed y adaptadores de storage
├── quiz.db               # SQLite database
├── Makefile              # Make commands
└── package.json          # NPM scripts
```

### Stack Tecnológico

**Backend:**
- Go 1.21+
- Gin (web framework)
- Arquitectura hexagonal: persistencia detrás del puerto `domain.Store`
- SQLite + GORM (adaptador por defecto, local)
- DynamoDB (adaptador para la nube/serverless, single-table sin GSIs)
- AWS CDK (infra/: Lambda + API Gateway + S3 + CloudFront + DynamoDB)

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
