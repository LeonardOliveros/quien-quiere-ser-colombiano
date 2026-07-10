<template>
  <div class="results-container">
    <div class="container py-5">
      <h1 class="text-center mb-5 text-golden">Resultados del Examen</h1>

      <div v-if="loading" class="text-center">
        <div class="spinner-border text-warning" role="status">
          <span class="visually-hidden">Cargando...</span>
        </div>
      </div>

      <div v-else-if="results">
        <!-- Score Card -->
        <div class="result-card">
          <div class="score-circle">
            <span class="score-big">{{ results.percentage.toFixed(1) }}%</span>
          </div>
          <div class="result-status" :class="{ 'passed': results.percentage >= 70, 'failed': results.percentage < 70 }">
            {{ results.percentage >= 70 ? '¡APROBADO!' : 'NO APROBADO' }}
          </div>
          <div class="mt-4">
            <p class="h4">{{ results.correct_answers }} de {{ results.total_questions }} correctas</p>
            <p class="text-muted">Tiempo total: {{ formatTime(results.time_taken) }}</p>
          </div>
        </div>

        <!-- Category Results -->
        <div class="result-card mt-4" v-if="results.category_scores && Object.keys(results.category_scores).length > 0">
          <h3 class="text-golden mb-4">Resultados por Categoría</h3>
          <div v-for="(categoryScore, category) in results.category_scores" :key="category" class="category-result">
            <h5>{{ categoryScore.category }}</h5>
            <div class="category-score-bar">
              <div class="category-score-fill" :style="{ width: categoryScore.percentage + '%' }">
                {{ categoryScore.percentage.toFixed(1) }}% ({{ categoryScore.correct_answers }}/{{ categoryScore.total_questions }})
              </div>
            </div>
          </div>
        </div>

        <!-- Incorrect Answers -->
        <div class="result-card mt-4" v-if="results.incorrect_answers && results.incorrect_answers.length > 0">
          <h3 class="text-golden mb-4">Respuestas Incorrectas</h3>
          <div class="incorrect-review">
            <div v-for="(item, index) in results.incorrect_answers" :key="index" class="review-item">
              <div class="review-question">{{ item.question.text }}</div>
              <div class="review-answer user-answer">
                <strong>Tu respuesta:</strong> {{ item.user_choice.text }}
              </div>
              <div class="review-answer correct-answer">
                <strong>Respuesta correcta:</strong> {{ item.correct_choice.text }}
              </div>
              <div class="explanation" v-if="item.explanation">
                <i class="fas fa-info-circle"></i> {{ item.explanation }}
              </div>
            </div>
          </div>
        </div>

        <!-- Flagged Questions -->
        <div class="result-card mt-4" v-if="results.flagged_questions && results.flagged_questions.length > 0">
          <h3 class="text-golden mb-4">Preguntas Marcadas</h3>
          <div class="flagged-review">
            <div v-for="(item, index) in results.flagged_questions" :key="index" class="review-item flagged">
              <div class="review-question">{{ item.text }}</div>
              <div class="badge bg-info">{{ item.category }}</div>
            </div>
          </div>
        </div>

        <!-- Recommendations -->
        <div class="result-card mt-4" v-if="results.recommendations && results.recommendations.length > 0">
          <h3 class="text-golden mb-4">Recomendaciones de Estudio</h3>
          <div class="recommendations">
            <div v-for="(rec, index) in results.recommendations" :key="index" class="recommendation-item">
              <p>{{ rec }}</p>
            </div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="text-center mt-5">
          <button class="btn btn-lg btn-primary me-3" @click="playAgain">
            <i class="fas fa-redo"></i> Jugar de Nuevo
          </button>
          <button class="btn btn-lg btn-outline-gold" @click="backToMenu">
            <i class="fas fa-home"></i> Volver al Menú
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useGameStore } from '@/stores/game'
import type { GameResults } from '@/types'

const router = useRouter()
const gameStore = useGameStore()

const loading = ref(true)
const results = ref<GameResults | null>(null)

onMounted(async () => {
  if (!gameStore.sessionId) {
    router.push('/')
    return
  }

  const result = await gameStore.loadResults()

  if (result.success && result.data) {
    results.value = result.data
  } else {
    alert(result.message || 'Error al cargar los resultados')
    router.push('/')
  }

  loading.value = false
})

function formatTime(seconds: number): string {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60

  if (hours > 0) {
    return `${hours}h ${minutes}m ${secs}s`
  } else if (minutes > 0) {
    return `${minutes}m ${secs}s`
  } else {
    return `${secs}s`
  }
}

function playAgain() {
  gameStore.resetGame()
  router.push('/')
}

function backToMenu() {
  gameStore.resetGame()
  router.push('/')
}
</script>

<style scoped>
.results-container {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  min-height: 100vh;
}

.text-golden {
  color: var(--gold-color);
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.result-card {
  background: rgba(255, 255, 255, 0.1);
  border: 2px solid var(--gold-color);
  border-radius: 20px;
  padding: 30px;
  margin-bottom: 20px;
  animation: fadeIn 0.5s ease-in;
}

.score-circle {
  width: 200px;
  height: 200px;
  border: 10px solid var(--gold-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 20px auto;
  position: relative;
  background: radial-gradient(circle, var(--accent-color), var(--secondary-color));
}

.score-big {
  font-size: 3rem;
  font-weight: bold;
  color: var(--gold-color);
}

.result-status {
  font-size: 2rem;
  font-weight: bold;
  margin-top: 20px;
}

.result-status.passed {
  color: var(--success-color);
}

.result-status.failed {
  color: var(--danger-color);
}

.category-result {
  background: rgba(255, 255, 255, 0.05);
  border-left: 4px solid var(--gold-color);
  padding: 15px;
  margin-bottom: 15px;
  border-radius: 5px;
}

.category-result h5 {
  color: var(--gold-color);
  margin-bottom: 10px;
}

.category-score-bar {
  background: rgba(0, 0, 0, 0.3);
  height: 30px;
  border-radius: 15px;
  overflow: hidden;
  position: relative;
}

.category-score-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--success-color), #34D399);
  transition: width 1s ease;
  display: flex;
  align-items: center;
  padding-left: 10px;
  font-weight: bold;
  color: white;
}

.incorrect-review, .flagged-review {
  max-height: 400px;
  overflow-y: auto;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  padding: 20px;
}

.review-item {
  background: rgba(255, 255, 255, 0.05);
  border-left: 4px solid var(--danger-color);
  padding: 15px;
  margin-bottom: 15px;
  border-radius: 5px;
}

.review-item.flagged {
  border-left-color: var(--warning-color);
}

.review-question {
  font-weight: bold;
  color: var(--gold-color);
  margin-bottom: 10px;
}

.review-answer {
  margin: 5px 0;
}

.user-answer {
  color: var(--danger-color);
}

.correct-answer {
  color: var(--success-color);
}

.explanation {
  background: rgba(255, 215, 0, 0.1);
  border-left: 3px solid var(--gold-color);
  padding: 10px;
  margin-top: 10px;
  font-style: italic;
}

.recommendations {
  background: rgba(0, 0, 0, 0.3);
  border-radius: 10px;
  padding: 20px;
}

.recommendation-item {
  background: linear-gradient(145deg, var(--accent-color), var(--secondary-color));
  border-left: 4px solid var(--warning-color);
  padding: 15px;
  margin-bottom: 15px;
  border-radius: 5px;
}

.recommendation-item h6 {
  color: var(--warning-color);
  margin-bottom: 10px;
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
  .score-circle {
    width: 150px;
    height: 150px;
  }

  .score-big {
    font-size: 2rem;
  }
}
</style>
