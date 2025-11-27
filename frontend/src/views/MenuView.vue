<template>
  <div class="menu-container">
    <div class="container py-5">
      <h1 class="display-3 text-center mb-5 text-golden">
        ¿Quién Quiere Ser Colombiano?
      </h1>

      <div class="game-modes">
        <!-- Resume Button (shown only when there's a paused game) -->
        <button v-if="hasPausedGame" class="btn-game btn-resume" @click="resumePausedGame">
          <i class="fas fa-play-circle"></i>
          <div>
            <strong>Reanudar Partida</strong>
            <br>
            <small>Continúa donde lo dejaste</small>
          </div>
        </button>

        <button class="btn-game" @click="startGame('practice')">
          <i class="fas fa-book-open"></i>
          <div>
            <strong>Modo Práctica</strong>
            <br>
            <small>Sin límite de tiempo</small>
          </div>
        </button>

        <button class="btn-game" @click="startGame('timed')">
          <i class="fas fa-clock"></i>
          <div>
            <strong>Contrarreloj</strong>
            <br>
            <small>3 horas para todas las preguntas</small>
          </div>
        </button>

        <button class="btn-game" @click="startGame('weak')">
          <i class="fas fa-chart-line"></i>
          <div>
            <strong>Áreas Débiles</strong>
            <br>
            <small>Enfócate en mejorar</small>
          </div>
        </button>

        <button class="btn-game" @click="showCategoryModal">
          <i class="fas fa-th-large"></i>
          <div>
            <strong>Por Categoría</strong>
            <br>
            <small>Selecciona una categoría</small>
          </div>
        </button>
      </div>

      <div class="text-center mt-5">
        <button class="btn btn-outline-gold me-2 mb-2" @click="showStatsModal">
          <i class="fas fa-chart-bar"></i> Mis Estadísticas
        </button>
        <button class="btn btn-outline-gold me-2 mb-2" @click="showHistoryModal">
          <i class="fas fa-history"></i> Historial
        </button>
        <button class="btn btn-outline-gold me-2 mb-2" @click="showRecommendationsModal">
          <i class="fas fa-lightbulb"></i> Recomendaciones
        </button>
        <button class="btn btn-outline-gold mb-2" @click="showQuestionCount">
          <i class="fas fa-database"></i> Base de Datos de Preguntas
        </button>
      </div>

      <div class="text-center mt-4">
        <button class="btn btn-danger" @click="logout">
          <i class="fas fa-sign-out-alt"></i> Cerrar Sesión
        </button>
      </div>
    </div>

    <!-- Stats Modal -->
    <StatsModal v-if="statsModalOpen" @close="statsModalOpen = false" />

    <!-- History Modal -->
    <HistoryModal v-if="historyModalOpen" @close="historyModalOpen = false" />

    <!-- Recommendations Modal -->
    <RecommendationsModal v-if="recommendationsModalOpen" @close="recommendationsModalOpen = false" />

    <!-- Category Modal -->
    <CategoryModal
      v-if="categoryModalOpen"
      @close="categoryModalOpen = false"
      @select="startCategoryGame"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useGameStore } from '@/stores/game'
import api from '@/services/api'
import StatsModal from '@/components/StatsModal.vue'
import HistoryModal from '@/components/HistoryModal.vue'
import RecommendationsModal from '@/components/RecommendationsModal.vue'
import CategoryModal from '@/components/CategoryModal.vue'

const router = useRouter()
const authStore = useAuthStore()
const gameStore = useGameStore()

const statsModalOpen = ref(false)
const historyModalOpen = ref(false)
const recommendationsModalOpen = ref(false)
const categoryModalOpen = ref(false)
const hasPausedGame = ref(false)
const pausedGameData = ref<any>(null)

// Check for paused game on mount
async function checkForPausedGame() {
  try {
    const data = await api.getAnyPausedGame()
    hasPausedGame.value = true
    pausedGameData.value = data
  } catch (error) {
    hasPausedGame.value = false
    pausedGameData.value = null
  }
}

// Call on component mount
checkForPausedGame()

function logout() {
  authStore.logout()
}

async function startGame(mode: string) {
  // Check if there's a paused game
  if (hasPausedGame.value && pausedGameData.value) {
    const choice = confirm(
      'Tienes una partida pausada.\n\n' +
      '¿Deseas cancelarla para iniciar una nueva?\n\n' +
      'Aceptar: Cancelar partida pausada\n' +
      'Cancelar: Mantener partida pausada'
    )

    if (!choice) {
      // User wants to keep the paused game
      return
    }

    // User wants to cancel the paused game - end it
    try {
      await api.endGame(pausedGameData.value.session_id.toString())
      hasPausedGame.value = false
      pausedGameData.value = null
    } catch (error) {
      console.error('Error ending paused game:', error)
    }
  }

  // Start new game
  let questionCount = 80
  let timeLimit = 0
  let categories: string[] = []
  let focusWeakAreas = false

  // For practice mode, get total question count from database
  if (mode === 'practice') {
    try {
      const countData = await api.getQuestionCount()
      questionCount = countData.total
    } catch (error) {
      console.error('Error getting question count:', error)
    }
  }

  // For timed mode, backend will force 80 questions
  if (mode === 'timed') {
    questionCount = 80 // Backend will enforce this
    timeLimit = 180 // 3 hours in minutes
  } else if (mode === 'weak') {
    focusWeakAreas = true
  }

  const result = await gameStore.startGame(
    mode,
    questionCount,
    timeLimit,
    categories,
    undefined,
    focusWeakAreas
  )

  if (result.success) {
    router.push('/game')
  } else {
    alert(result.message || 'Error al iniciar el juego')
  }
}

async function resumePausedGame() {
  if (!pausedGameData.value) return

  const resumeResult = await gameStore.resumeGame(pausedGameData.value)
  if (resumeResult.success) {
    router.push('/game')
  } else {
    alert(resumeResult.message || 'Error al reanudar la partida')
  }
}

function showCategoryModal() {
  categoryModalOpen.value = true
}

async function startCategoryGame(category: string) {
  const result = await gameStore.startGame(
    'category',
    80,
    0,
    [category],
    undefined,
    false
  )

  if (result.success) {
    router.push('/game')
  } else {
    alert(result.message || 'Error al iniciar el juego')
  }
}

function showStatsModal() {
  statsModalOpen.value = true
}

function showHistoryModal() {
  historyModalOpen.value = true
}

function showRecommendationsModal() {
  recommendationsModalOpen.value = true
}

async function showQuestionCount() {
  try {
    const data = await api.getQuestionCount()
    let message = `Total de preguntas: ${data.total}\n\nPor categoría:\n`
    for (const [category, count] of Object.entries(data.by_category)) {
      message += `${category}: ${count}\n`
    }
    alert(message)
  } catch (error) {
    alert('Error al cargar el conteo de preguntas')
  }
}
</script>

<style scoped>
.menu-container {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  min-height: 100vh;
  animation: fadeIn 0.5s ease-in;
}

.text-golden {
  color: var(--gold-color);
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
  font-weight: bold;
  letter-spacing: 2px;
}

.game-modes {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
  margin-bottom: 40px;
}

.btn-game {
  background: linear-gradient(145deg, var(--accent-color), var(--secondary-color));
  color: var(--text-light);
  border: 2px solid var(--gold-color);
  padding: 20px 30px;
  border-radius: 15px;
  transition: all 0.3s ease;
  min-width: 250px;
  position: relative;
  overflow: hidden;
  cursor: pointer;
}

.btn-game:before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(45deg, var(--gold-color), transparent);
  z-index: -1;
  opacity: 0;
  transition: opacity 0.3s ease;
  border-radius: 15px;
}

.btn-game:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(255, 215, 0, 0.3);
  color: var(--gold-color);
}

.btn-game:hover:before {
  opacity: 1;
}

.btn-game i {
  font-size: 2rem;
  display: block;
  margin-bottom: 10px;
}

.btn-game small {
  font-size: 0.9rem;
  opacity: 0.9;
  margin-top: 5px;
}

.btn-resume {
  background: linear-gradient(145deg, var(--accent-teal), #1fa89a);
  border-color: var(--accent-teal);
  animation: pulse 2s ease-in-out infinite;
}

.btn-resume:hover {
  background: linear-gradient(145deg, #3FFFDD, var(--accent-teal));
  box-shadow: 0 10px 30px rgba(43, 217, 185, 0.5);
}

@keyframes pulse {
  0%, 100% {
    box-shadow: 0 0 0 0 rgba(43, 217, 185, 0.7);
  }
  50% {
    box-shadow: 0 0 0 10px rgba(43, 217, 185, 0);
  }
}

.btn-outline-gold {
  background: transparent;
  border: 2px solid var(--gold-color);
  color: var(--gold-color);
  padding: 10px 20px;
  border-radius: 10px;
  transition: all 0.3s ease;
}

.btn-outline-gold:hover {
  background: var(--gold-color);
  color: var(--primary-color);
  transform: translateY(-2px);
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
  .btn-game {
    min-width: 200px;
    padding: 15px 20px;
  }
}
</style>
