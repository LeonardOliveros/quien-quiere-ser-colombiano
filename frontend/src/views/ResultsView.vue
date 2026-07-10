<template>
  <div class="results-container">
    <div class="container py-5 results-content">
      <h1 class="text-center mb-4 text-golden">Resultados del Examen</h1>

      <div v-if="loading" class="text-center">
        <div class="spinner-border text-warning" role="status">
          <span class="visually-hidden">Cargando...</span>
        </div>
      </div>

      <div v-else-if="results">
        <!-- Score Hero -->
        <div class="result-card hero-card">
          <div class="hero-score">
            <svg class="score-ring" viewBox="0 0 120 120" role="img"
                 :aria-label="`Puntaje ${Math.round(results.percentage)}%`">
              <circle class="ring-track" cx="60" cy="60" r="52" />
              <circle
                class="ring-fill"
                :class="passed ? 'ring-pass' : 'ring-fail'"
                cx="60" cy="60" r="52"
                :stroke-dasharray="`${ringLength} ${RING_CIRCUMFERENCE}`"
                transform="rotate(-90 60 60)"
              />
            </svg>
            <div class="score-ring-label">
              <span class="score-big">{{ Math.round(results.percentage) }}%</span>
              <span class="score-sub">de acierto</span>
            </div>
          </div>

          <div class="hero-details">
            <div class="status-chip" :class="passed ? 'chip-pass' : 'chip-fail'">
              <i :class="passed ? 'fas fa-check-circle' : 'fas fa-times-circle'"></i>
              {{ passed ? '¡Aprobado!' : 'No aprobado' }}
            </div>
            <p class="hero-note">Necesitas un 70% de aciertos para aprobar el examen</p>

            <div class="stat-tiles">
              <div class="stat-tile">
                <span class="stat-value stat-correct">{{ results.correct_answers }}</span>
                <span class="stat-label">Correctas</span>
              </div>
              <div class="stat-tile">
                <span class="stat-value stat-incorrect">{{ incorrectCount }}</span>
                <span class="stat-label">Incorrectas</span>
              </div>
              <div class="stat-tile">
                <span class="stat-value">{{ results.score }}</span>
                <span class="stat-label">Puntos</span>
              </div>
              <div class="stat-tile">
                <span class="stat-value">{{ formatTime(results.time_taken) }}</span>
                <span class="stat-label">Tiempo total</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Category Results -->
        <div class="result-card section-card" v-if="categoryList.length > 0">
          <h3 class="section-title">Resultados por Categoría</h3>
          <div v-for="cat in categoryList" :key="cat.category" class="category-row">
            <div class="category-head">
              <span class="category-name">{{ cat.category }}</span>
              <span class="category-meta">
                <span class="category-frac">{{ cat.correct_answers }}/{{ cat.total_questions }} correctas</span>
                <span class="category-pct">{{ Math.round(cat.percentage) }}%</span>
                <span class="pill" :class="cat.passed ? 'pill-pass' : 'pill-fail'">
                  <i :class="cat.passed ? 'fas fa-check' : 'fas fa-times'"></i>
                  {{ cat.passed ? 'Superada' : 'Por reforzar' }}
                </span>
              </span>
            </div>
            <div class="category-bar">
              <div
                class="category-fill"
                :class="cat.passed ? 'fill-pass' : 'fill-fail'"
                :style="{ width: Math.max(cat.percentage, 0) + '%' }"
              ></div>
            </div>
          </div>
        </div>

        <!-- Incorrect Answers -->
        <div class="result-card section-card" v-if="results.incorrect_answers && results.incorrect_answers.length > 0">
          <h3 class="section-title">
            Respuestas Incorrectas
            <span class="count-badge">{{ results.incorrect_answers.length }}</span>
          </h3>
          <div class="review-list">
            <article v-for="(item, index) in results.incorrect_answers" :key="index" class="review-card">
              <p class="review-question">{{ item.question.text }}</p>
              <div class="answer-line answer-wrong">
                <i class="fas fa-times" aria-hidden="true"></i>
                <p><span class="answer-tag">Tu respuesta</span>{{ item.user_choice.text }}</p>
              </div>
              <div class="answer-line answer-right">
                <i class="fas fa-check" aria-hidden="true"></i>
                <p><span class="answer-tag">Correcta</span>{{ item.correct_choice.text }}</p>
              </div>
              <p class="review-explanation" v-if="item.explanation">
                <i class="fas fa-info-circle" aria-hidden="true"></i> {{ item.explanation }}
              </p>
            </article>
          </div>
        </div>

        <!-- Flagged Questions -->
        <div class="result-card section-card" v-if="results.flagged_questions && results.flagged_questions.length > 0">
          <h3 class="section-title">
            Preguntas Marcadas
            <span class="count-badge">{{ results.flagged_questions.length }}</span>
          </h3>
          <div class="review-list">
            <div v-for="(item, index) in results.flagged_questions" :key="index" class="flagged-row">
              <i class="fas fa-flag" aria-hidden="true"></i>
              <p class="flagged-text">{{ item.text }}</p>
              <span class="badge bg-info">{{ item.category }}</span>
            </div>
          </div>
        </div>

        <!-- Recommendations -->
        <div class="result-card section-card" v-if="results.recommendations && results.recommendations.length > 0">
          <h3 class="section-title">Recomendaciones de Estudio</h3>
          <div class="review-list">
            <div v-for="(rec, index) in results.recommendations" :key="index" class="recommendation-row">
              <i class="fas fa-book-open" aria-hidden="true"></i>
              <p>{{ rec }}</p>
            </div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="actions mt-4">
          <button class="btn btn-lg btn-primary" @click="playAgain">
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
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useGameStore } from '@/stores/game'
import type { GameResults } from '@/types'

const router = useRouter()
const gameStore = useGameStore()

const loading = ref(true)
const results = ref<GameResults | null>(null)

// SVG ring geometry: 2 * PI * r (r = 52)
const RING_CIRCUMFERENCE = 326.7

const passed = computed(() => (results.value?.percentage ?? 0) >= 70)

const ringLength = computed(() => {
  const pct = Math.min(Math.max(results.value?.percentage ?? 0, 0), 100)
  return (pct / 100) * RING_CIRCUMFERENCE
})

const incorrectCount = computed(() => {
  if (!results.value) return 0
  return results.value.total_questions - results.value.correct_answers
})

// Fixed order so categories always appear in the same position
const categoryList = computed(() => {
  if (!results.value?.category_scores) return []
  return Object.values(results.value.category_scores)
    .sort((a, b) => a.category.localeCompare(b.category))
})

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

.results-content {
  max-width: 860px;
}

.text-golden {
  color: var(--gold-color);
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.result-card {
  background: var(--bg-card);
  border: 1px solid rgba(255, 205, 0, 0.35);
  border-radius: 16px;
  padding: 28px;
  margin-bottom: 20px;
  text-align: left;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.35);
  animation: fadeIn 0.5s ease-in;
}

/* ---------- Hero ---------- */
.hero-card {
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 32px;
  align-items: center;
}

.hero-score {
  position: relative;
  width: 190px;
  height: 190px;
  margin: 0 auto;
}

.score-ring {
  width: 100%;
  height: 100%;
}

.ring-track {
  fill: none;
  stroke: rgba(255, 255, 255, 0.12);
  stroke-width: 10;
}

.ring-fill {
  fill: none;
  stroke-width: 10;
  stroke-linecap: round;
  transition: stroke-dasharray 1s ease;
}

.ring-pass { stroke: var(--emerald); }
.ring-fail { stroke: var(--flag-red-light); }

.score-ring-label {
  position: absolute;
  inset: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2px;
}

.score-big {
  font-size: 2.6rem;
  font-weight: 800;
  line-height: 1;
  color: var(--text-main);
}

.score-sub {
  font-size: 0.85rem;
  color: rgba(253, 249, 238, 0.6);
}

.hero-details {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.status-chip {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  align-self: flex-start;
  padding: 6px 16px;
  border-radius: 999px;
  font-size: 1.1rem;
  font-weight: 700;
}

.chip-pass {
  color: var(--emerald);
  background: rgba(16, 185, 129, 0.14);
  border: 1px solid rgba(16, 185, 129, 0.5);
}

.chip-fail {
  color: var(--flag-red-light);
  background: rgba(232, 69, 91, 0.14);
  border: 1px solid rgba(232, 69, 91, 0.5);
}

.hero-note {
  margin: 0;
  font-size: 0.9rem;
  color: rgba(253, 249, 238, 0.6);
}

.stat-tiles {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
  margin-top: 8px;
}

.stat-tile {
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 12px;
  padding: 12px 14px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.stat-value {
  font-size: 1.4rem;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.2;
}

.stat-correct { color: var(--emerald); }
.stat-incorrect { color: var(--flag-red-light); }

.stat-label {
  font-size: 0.78rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: rgba(253, 249, 238, 0.55);
}

/* ---------- Sections ---------- */
.section-title {
  display: flex;
  align-items: center;
  gap: 10px;
  color: var(--gold-color);
  font-size: 1.25rem;
  margin-bottom: 20px;
}

.count-badge {
  font-size: 0.8rem;
  font-weight: 700;
  color: var(--text-main);
  background: rgba(255, 255, 255, 0.12);
  border-radius: 999px;
  padding: 2px 10px;
}

/* ---------- Category bars ---------- */
.category-row {
  padding: 12px 0;
}

.category-row + .category-row {
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

.category-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 12px;
  flex-wrap: wrap;
  margin-bottom: 8px;
}

.category-name {
  font-weight: 700;
  letter-spacing: 0.5px;
  color: var(--text-main);
}

.category-meta {
  display: inline-flex;
  align-items: center;
  gap: 12px;
}

.category-frac {
  font-size: 0.85rem;
  color: rgba(253, 249, 238, 0.6);
}

.category-pct {
  font-weight: 700;
  font-variant-numeric: tabular-nums;
  color: var(--text-main);
  min-width: 3ch;
  text-align: right;
}

.pill {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 0.75rem;
  font-weight: 700;
  padding: 3px 10px;
  border-radius: 999px;
}

.pill-pass {
  color: var(--emerald);
  background: rgba(16, 185, 129, 0.14);
}

.pill-fail {
  color: var(--flag-red-light);
  background: rgba(232, 69, 91, 0.14);
}

.category-bar {
  background: rgba(0, 0, 0, 0.35);
  height: 10px;
  border-radius: 5px;
  overflow: hidden;
}

.category-fill {
  height: 100%;
  border-radius: 5px;
  transition: width 1s ease;
}

.fill-pass { background: var(--emerald); }
.fill-fail { background: var(--flag-red-light); }

/* ---------- Review lists ---------- */
.review-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.review-card {
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 12px;
  padding: 16px 18px;
}

.review-question {
  font-weight: 600;
  color: var(--text-main);
  margin: 0 0 12px;
  line-height: 1.45;
}

.answer-line {
  display: flex;
  align-items: baseline;
  gap: 10px;
  margin: 6px 0;
}

.answer-line p {
  margin: 0;
  color: var(--text-main);
}

.answer-line i {
  width: 16px;
  text-align: center;
}

.answer-wrong i { color: var(--flag-red-light); }
.answer-right i { color: var(--emerald); }

.answer-tag {
  display: inline-block;
  font-size: 0.72rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: rgba(253, 249, 238, 0.55);
  margin-right: 8px;
}

.review-explanation {
  background: rgba(255, 205, 0, 0.08);
  border-left: 3px solid var(--gold-color);
  border-radius: 0 8px 8px 0;
  padding: 10px 14px;
  margin: 12px 0 0;
  font-size: 0.92rem;
  color: rgba(253, 249, 238, 0.85);
}

.review-explanation i {
  color: var(--gold-color);
  margin-right: 6px;
}

/* ---------- Flagged & recommendations ---------- */
.flagged-row,
.recommendation-row {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 12px;
  padding: 14px 18px;
}

.flagged-row i { color: var(--gold-color); margin-top: 3px; }
.recommendation-row i { color: var(--gold-color); margin-top: 3px; }

.flagged-text,
.recommendation-row p {
  flex: 1;
  margin: 0;
  color: var(--text-main);
}

.flagged-row .badge {
  flex-shrink: 0;
  align-self: center;
}

/* ---------- Actions ---------- */
.actions {
  display: flex;
  justify-content: center;
  gap: 16px;
  flex-wrap: wrap;
}

.btn-outline-gold {
  background: transparent;
  border: 2px solid var(--gold-color);
  color: var(--gold-color);
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

/* ---------- Responsive ---------- */
@media (max-width: 720px) {
  .hero-card {
    grid-template-columns: 1fr;
    gap: 20px;
    text-align: center;
  }

  .hero-score {
    width: 160px;
    height: 160px;
  }

  .status-chip {
    align-self: center;
  }

  .stat-tiles {
    grid-template-columns: repeat(2, 1fr);
  }

  .category-head {
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }
}
</style>
