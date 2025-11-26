# Test de Pause/Resume - Guía de Prueba Manual

## Estado Actual de la Base de Datos
```bash
sqlite3 quiz.db "SELECT id, user_id, mode, status, total_questions, correct_answers, (SELECT COUNT(*) FROM game_answers WHERE game_session_id = game_sessions.id) as answered FROM game_sessions WHERE id >= 6 ORDER BY id DESC;"
```

## Pasos para Probar Pause/Resume

### Preparación
1. Asegúrate de estar logueado como el usuario `leonardo` (user_id: 1)
2. Limpia sesiones antiguas si es necesario:
   ```bash
   make clean-sessions
   ```

### Test 1: Crear y Pausar una Sesión
1. Inicia la aplicación: `./quiz-app`
2. Login con usuario `leonardo`
3. Selecciona "Modo Práctica"
4. Responde **3-5 preguntas** (mezcla de correctas e incorrectas)
5. Marca **2 preguntas** con el botón "Marcar"
6. Click en botón **"Pausar"**
7. Confirma pausar
8. Deberías volver al menú principal

### Test 2: Verificar Sesión Pausada en DB
```bash
sqlite3 quiz.db "SELECT id, mode, status, total_questions, correct_answers, (SELECT COUNT(*) FROM game_answers WHERE game_session_id = game_sessions.id) as answered, (SELECT COUNT(*) FROM game_answers WHERE game_session_id = game_sessions.id AND is_flagged = 1) as flagged FROM game_sessions WHERE status = 'PAUSED' ORDER BY id DESC LIMIT 1;"
```

Deberías ver algo como:
```
<session_id>|practice|PAUSED|753|<correct_count>|<answered_count>|2
```

### Test 3: Reanudar Sesión
1. Click en "Modo Práctica" nuevamente
2. Deberías ver un confirm dialog que dice:
   ```
   Tienes una partida pausada (X/753 preguntas respondidas).

   ¿Deseas reanudarla?

   Aceptar: Reanudar partida
   Cancelar: Nueva partida
   ```
3. Click en **"Aceptar"**
4. Deberías ver:
   - La pregunta número (X+1) donde X es el número de preguntas respondidas
   - El contador de correctas/incorrectas debe coincidir con los valores previos
   - Las preguntas marcadas deben seguir marcadas

### Test 4: Verificar Estado en DB
```bash
sqlite3 quiz.db "SELECT id, mode, status FROM game_sessions ORDER BY id DESC LIMIT 3;"
```

Deberías ver que:
- La sesión reanudada cambió de PAUSED a ACTIVE
- NO se creó una nueva sesión
- El session_id es el mismo que antes de pausar

### Test 5: Completar Sesión Reanudada
1. Responde algunas preguntas más
2. Click en "Salir" o completa todas las preguntas
3. Verifica en resultados que los contadores son correctos

### Test 6: Nueva Partida en Lugar de Reanudar
1. Repite Test 1 (crear sesión pausada)
2. Click en "Modo Práctica"
3. En el dialog, click en **"Cancelar"** (Nueva partida)
4. Verifica que comienza una nueva sesión desde la pregunta 1
5. Verifica en DB:
   ```bash
   sqlite3 quiz.db "SELECT id, mode, status FROM game_sessions ORDER BY id DESC LIMIT 2;"
   ```
   - La sesión pausada anterior debe estar COMPLETED
   - Debe haber una nueva sesión ACTIVE

## Logs Esperados

### Al Reanudar (logs en consola del servidor):
```
GET "/api/game/paused/PRACTICE" → 200
GET "/api/game/<session_id>/question" → 200 (con status cambiando a ACTIVE)
```

### Al Crear Nueva (logs en consola del servidor):
```
GET "/api/game/paused/PRACTICE" → 200
POST "/api/game/<old_session_id>/end" → 200
POST "/api/game/start" → 200
GET "/api/game/<new_session_id>/question" → 200
```

## Problemas Conocidos Corregidos

✅ **CORREGIDO**: Frontend obtenía 753 preguntas para modo TIMED en lugar de 80
✅ **CORREGIDO**: Sesiones PAUSED antiguas se acumulaban
✅ **CORREGIDO**: Error NaN al finalizar juego sin responder preguntas
✅ **CORREGIDO**: No se finalizaba sesión pausada al crear nueva partida

## Si Algo Falla

### Session ID no se mantiene al reanudar
- Verifica que `resumeGame` en `game.ts` está usando `pausedData.session_id`
- Verifica logs del servidor para confirmar que GET `/api/game/<session_id>/question` usa el session_id correcto

### No se detecta sesión pausada
- Verifica que estás logueado con el mismo usuario
- Verifica que el modo coincide (PRACTICE, TIMED, etc.)
- Verifica en DB que existe una sesión PAUSED para ese user_id y mode

### Contadores no coinciden al reanudar
- Verifica que `resumeGame` restaura: correctAnswers, incorrectAnswers, flaggedCount, flaggedQuestions
- Verifica que backend devuelve estos valores en `/api/game/paused/<mode>`
