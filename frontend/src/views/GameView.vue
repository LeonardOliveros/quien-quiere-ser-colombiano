<template>
  <div class="game-container">
    <div class="container-fluid px-4 py-3">
      <!-- Game Header -->
      <div class="game-header">
        <div class="row align-items-center">
          <div class="col-md-3">
            <div v-if="gameStore.timeLimit > 0" class="timer-display">
              <i class="fas fa-clock"></i> Restante: {{ formatTime(timeLeft) }}
            </div>
            <div v-else class="timer-display">
              <i class="fas fa-stopwatch"></i> Tiempo: {{ formatElapsedTime(elapsedTime) }}
            </div>
          </div>
          <div class="col-md-6 text-center">
            <div class="question-counter">
              Pregunta {{ gameStore.questionNumber }} de {{ gameStore.totalQuestions }}
            </div>
          </div>
          <div class="col-md-3 text-end">
            <button
              class="btn btn-sm me-2"
              :class="isFlagged ? 'btn-warning' : 'btn-outline-warning'"
              @click="flagCurrentQuestion"
            >
              <i class="fas fa-flag"></i> {{ isFlagged ? 'Marcada' : 'Marcar' }}
            </button>
            <button class="btn btn-sm btn-info me-2" @click="confirmPause">
              <i class="fas fa-pause"></i> Pausar
            </button>
            <button class="btn btn-sm btn-danger" @click="confirmQuit">
              <i class="fas fa-times"></i> Salir
            </button>
          </div>
        </div>
        <div class="row mt-2">
          <div class="col-12 text-center category-badge">
            <span class="badge bg-info" v-if="gameStore.currentQuestion">
              {{ gameStore.currentQuestion.category }}
            </span>
          </div>
        </div>
      </div>

      <!-- Question Box -->
      <div v-if="gameStore.currentQuestion" class="question-box">
        <p class="question-text">{{ gameStore.currentQuestion.text }}</p>
      </div>
      <div v-else class="question-box" aria-busy="true">
        <span class="visually-hidden">Cargando pregunta...</span>
        <SkeletonBlock width="90%" height="1.3rem" radius="4px" class="mx-auto" />
        <SkeletonBlock width="65%" height="1.3rem" radius="4px" class="mx-auto mt-3" />
      </div>

      <!-- Answers Grid -->
      <div v-if="gameStore.currentQuestion" class="answers-grid">
        <div class="row g-3">
          <div
            v-for="(choice, index) in visibleChoices"
            :key="choice.id"
            class="col-md-6"
          >
            <button
              class="answer-btn"
              :class="{
                'disabled': answering || showingAnswer,
                'correct': showingAnswer && choice.id === correctChoiceId,
                'incorrect': showingAnswer && choice.id === selectedChoiceId && !answerCorrect
              }"
              @click="selectAnswer(choice.id)"
              :disabled="answering || showingAnswer"
            >
              <span class="answer-letter">{{ String.fromCharCode(65 + index) }}</span>
              <span class="answer-text">{{ choice.text }}</span>
            </button>
          </div>
        </div>
      </div>
      <div v-else class="answers-grid" aria-busy="true">
        <div class="row g-3">
          <div v-for="n in 4" :key="n" class="col-md-6">
            <SkeletonBlock height="68px" radius="15px" />
          </div>
        </div>
      </div>

      <!-- Explanation shown below answers after selection -->
      <div v-if="showingAnswer && answerExplanation" class="explanation-box mt-3">
        <strong><i class="fas fa-info-circle"></i> Explicación:</strong> {{ answerExplanation }}
      </div>

      <!-- Lifelines -->
      <div class="lifelines text-center">
        <button
          class="btn-lifeline"
          :class="{ 'used': gameStore.fiftyFiftyRemaining === 0 || fiftyFiftyUsedOnCurrentQuestion }"
          @click="useFiftyFifty"
          :disabled="gameStore.fiftyFiftyRemaining === 0 || fiftyFiftyUsedOnCurrentQuestion || answering || showingAnswer"
        >
          <i class="fas fa-divide"></i> 50:50 ({{ gameStore.fiftyFiftyRemaining }})
        </button>
        <button
          class="btn-lifeline"
          :class="{ 'used': gameStore.autosolveRemaining === 0 }"
          @click="useAutosolve"
          :disabled="gameStore.autosolveRemaining === 0 || answering || showingAnswer"
        >
          <i class="fas fa-check-circle"></i> Resolver ({{ gameStore.autosolveRemaining }})
        </button>
        <button
          class="btn-lifeline"
          :class="{ 'used': gameStore.skipsRemaining === 0 }"
          @click="skipQuestion"
          :disabled="gameStore.skipsRemaining === 0 || answering || showingAnswer"
        >
          <i class="fas fa-forward"></i> Saltar ({{ gameStore.skipsRemaining }})
        </button>
      </div>

      <!-- Progress Indicator -->
      <div class="progress-indicator">
        <div class="progress">
          <div
            class="progress-bar"
            :style="{ width: progressPercentage + '%' }"
            role="progressbar"
          >
            {{ progressPercentage }}%
          </div>
        </div>

        <div class="score-display mt-3">
          <span class="badge bg-success">
            <i class="fas fa-check"></i> Correctas: {{ gameStore.correctAnswers }}
          </span>
          <span class="badge bg-danger">
            <i class="fas fa-times"></i> Incorrectas: {{ gameStore.incorrectAnswers }}
          </span>
          <span class="badge bg-warning">
            <i class="fas fa-flag"></i> Marcadas: {{ gameStore.flaggedCount }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useGameStore } from '@/stores/game'
import api from '@/services/api'
import SkeletonBlock from '@/components/SkeletonBlock.vue'

const router = useRouter()
const gameStore = useGameStore()

const answering = ref(false)
const showingAnswer = ref(false)
const answerCorrect = ref(false)
const answerExplanation = ref('')
const correctAnswerText = ref('')
const correctChoiceId = ref<number | null>(null)
const selectedChoiceId = ref<number | null>(null)
const removedChoices = ref<Set<number>>(new Set())
const fiftyFiftyUsedOnCurrentQuestion = ref(false)
const timeLeft = ref(0)
const elapsedTime = ref(0)
const isFlagged = ref(false)
let timerInterval: number | null = null

const visibleChoices = computed(() => {
  if (!gameStore.currentQuestion) return []
  return gameStore.currentQuestion.choices.filter(
    (choice) => !removedChoices.value.has(choice.id)
  )
})

const progressPercentage = computed(() => {
  if (gameStore.totalQuestions === 0) return 0
  return Math.round((gameStore.questionNumber / gameStore.totalQuestions) * 100)
})

onMounted(async () => {
  if (!gameStore.sessionId) {
    router.push('/')
    return
  }

  await loadQuestion()

  // Start timer - use time_elapsed from backend to continue where we left off
  if (gameStore.timeLimit > 0) {
    // Countdown timer for timed mode (timeLimit is in seconds)
    timeLeft.value = Math.max(0, gameStore.timeLimit - gameStore.timeElapsed)
    startTimer()
  } else {
    // Elapsed time counter for practice mode - continue from where we left off
    elapsedTime.value = gameStore.timeElapsed
    startElapsedTimer()
  }
})

onUnmounted(() => {
  if (timerInterval) {
    clearInterval(timerInterval)
  }
})

async function loadQuestion() {
  const result = await gameStore.loadNextQuestion()

  if (!result.success) {
    if (result.noMoreQuestions) {
      await finishGame()
    } else {
      alert(result.message || 'Error al cargar la pregunta')
    }
    return
  }

  // Re-sync the countdown with the server on every question so the local
  // timer never drifts (e.g. throttled intervals in background tabs)
  if (gameStore.timeLimit > 0 && gameStore.timeRemaining >= 0) {
    timeLeft.value = gameStore.timeRemaining
  }

  // Reset question state
  showingAnswer.value = false
  removedChoices.value.clear()
  fiftyFiftyUsedOnCurrentQuestion.value = false
  selectedChoiceId.value = null
  correctChoiceId.value = null

  // Check if current question is already flagged
  if (gameStore.currentQuestion) {
    isFlagged.value = gameStore.isQuestionFlagged(gameStore.currentQuestion.id)
  }
}

async function selectAnswer(choiceId: number) {
  if (answering.value || showingAnswer.value) return

  answering.value = true
  selectedChoiceId.value = choiceId

  const result = await gameStore.submitAnswer(choiceId)

  if (result.success) {
    answerCorrect.value = result.correct!
    answerExplanation.value = result.explanation!
    correctChoiceId.value = result.correctChoiceId!

    if (!result.correct && gameStore.currentQuestion) {
      const correctChoice = gameStore.currentQuestion.choices.find(
        (c) => c.id === result.correctChoiceId
      )
      correctAnswerText.value = correctChoice?.text || ''
    }

    showingAnswer.value = true

    // Auto-advance to next question after 3 seconds
    setTimeout(async () => {
      await loadQuestion()
    }, 3000)
  } else if (result.timeUp) {
    // The server ended the session (time limit) — go to results
    await finishGame()
  } else {
    alert(result.message || 'Error al enviar la respuesta')
  }

  answering.value = false
}

async function flagCurrentQuestion() {
  const result = await gameStore.flagQuestion()
  if (result.success) {
    // Update the visual state based on the store's response
    isFlagged.value = result.isFlagged!
  } else {
    alert(result.message || 'Error al marcar la pregunta')
  }
}

async function useFiftyFifty() {
  if (answering.value || showingAnswer.value) return
  if (gameStore.useFiftyFifty() && gameStore.currentQuestion && gameStore.sessionId) {
    try {
      // Call backend to get which incorrect choices to remove
      const result = await api.useFiftyFifty(
        gameStore.sessionId,
        gameStore.currentQuestion.id
      )

      // Remove the choices returned by the backend
      result.remove_choice_ids.forEach((choiceId: number) => {
        removedChoices.value.add(choiceId)
      })

      // Mark that 50:50 was used on this question
      fiftyFiftyUsedOnCurrentQuestion.value = true
    } catch (error) {
      console.error('Error using 50-50:', error)
      alert('Error al usar el comodín 50-50')
    }
  }
}

async function useAutosolve() {
  if (answering.value || showingAnswer.value) return
  if (gameStore.useAutosolve() && gameStore.currentQuestion && gameStore.sessionId) {
    try {
      // Call backend to get the correct choice
      const result = await api.useAutosolve(
        gameStore.sessionId,
        gameStore.currentQuestion.id
      )

      // Auto-select the correct answer
      await selectAnswer(result.correct_choice_id)
    } catch (error) {
      console.error('Error using autosolve:', error)
      alert('Error al usar el comodín de resolver')
    }
  }
}

async function skipQuestion() {
  if (answering.value || showingAnswer.value) return
  if (gameStore.useSkip()) {
    await loadQuestion()
  }
}

function confirmPause() {
  if (confirm('¿Deseas pausar la partida? Podrás reanudarla más tarde.')) {
    pauseGame()
  }
}

async function pauseGame() {
  const result = await gameStore.pauseGame()
  if (result.success) {
    alert('Partida pausada. Podrás reanudarla desde el menú principal.')
    router.push('/')
  } else {
    alert(result.message || 'Error al pausar la partida')
  }
}

function confirmQuit() {
  if (confirm('¿Estás seguro de que quieres salir del juego?')) {
    finishGame()
  }
}

async function finishGame() {
  await gameStore.endGame()
  router.push('/results')
}

function startTimer() {
  timerInterval = window.setInterval(() => {
    if (timeLeft.value > 0) {
      timeLeft.value--
    } else {
      clearInterval(timerInterval!)
      finishGame()
    }
  }, 1000)
}

function formatTime(seconds: number): string {
  const total = Math.max(0, seconds)
  const hours = Math.floor(total / 3600)
  const minutes = Math.floor((total % 3600) / 60)
  const secs = total % 60
  return `${hours}:${String(minutes).padStart(2, '0')}:${String(secs).padStart(2, '0')}`
}

function startElapsedTimer() {
  timerInterval = window.setInterval(() => {
    elapsedTime.value++
  }, 1000)
}

function formatElapsedTime(seconds: number): string {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  return `${hours}:${String(minutes).padStart(2, '0')}:${String(secs).padStart(2, '0')}`
}
</script>

<style scoped>
.game-container {
  background: radial-gradient(ellipse at center, var(--secondary-color) 0%, var(--primary-color) 100%);
  min-height: 100vh;
  padding-bottom: 40px;
}

.game-header {
  background: rgba(0, 0, 0, 0.3);
  padding: 15px;
  border-radius: 10px;
  margin-bottom: 30px;
}

.timer-display {
  font-size: 1.25rem;
  color: var(--warning-color);
  font-weight: bold;
  /* Digits with a fixed width so the header doesn't jiggle every second */
  font-variant-numeric: tabular-nums;
  white-space: nowrap;
}

.question-counter {
  font-size: 1.2rem;
  color: var(--gold-color);
  font-weight: 600;
}

.category-badge .badge {
  font-size: 1rem;
  padding: 8px 20px;
  text-transform: uppercase;
}

.question-box {
  background: linear-gradient(145deg, var(--accent-color), var(--secondary-color));
  border: 3px solid var(--gold-color);
  border-radius: 20px;
  padding: 40px;
  margin: 30px auto;
  max-width: 900px;
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.5);
  position: relative;
}

.question-box::before {
  content: '';
  position: absolute;
  top: -5px;
  left: -5px;
  right: -5px;
  bottom: -5px;
  background: linear-gradient(45deg, var(--gold-color), transparent, var(--gold-color));
  z-index: -1;
  border-radius: 20px;
  animation: glow 2s ease-in-out infinite;
}

@keyframes glow {
  0%, 100% { opacity: 0.5; }
  50% { opacity: 1; }
}

.question-text {
  font-size: 1.8rem;
  text-align: center;
  color: var(--text-light);
  margin: 0;
  line-height: 1.4;
}

.hint-box {
  background: rgba(255, 193, 7, 0.2);
  border: 2px solid var(--warning-color);
  border-radius: 10px;
  padding: 20px;
  margin: 20px auto;
  max-width: 900px;
  text-align: center;
  color: var(--warning-color);
}

.answers-grid {
  max-width: 900px;
  margin: 0 auto;
}

.answer-btn {
  width: 100%;
  background: linear-gradient(145deg, #0A2A6B, #14418F);
  border: 2px solid var(--gold-color);
  color: var(--text-light);
  padding: 20px;
  border-radius: 15px;
  font-size: 1.1rem;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.answer-btn:hover:not(.disabled) {
  background: linear-gradient(145deg, #14418F, #0A2A6B);
  transform: translateX(5px);
  box-shadow: 0 5px 20px var(--glow-yellow);
}

.answer-btn.disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.answer-letter {
  width: 40px;
  height: 40px;
  background: var(--gold-color);
  color: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  margin-right: 20px;
  font-size: 1.2rem;
  flex-shrink: 0;
}

.answer-text {
  flex: 1;
  text-align: left;
}

.answer-feedback {
  max-width: 900px;
  margin: 20px auto;
}

.feedback-message {
  animation: fadeIn 0.5s ease;
}

.feedback-message h3 {
  font-weight: bold;
  font-size: 1.8rem;
  margin: 10px 0;
}

.explanation-box {
  background: rgba(255, 215, 0, 0.1);
  border-left: 4px solid var(--gold-color);
  padding: 20px;
  text-align: left;
  border-radius: 5px;
  margin: 20px auto;
  max-width: 900px;
  animation: slideIn 0.5s ease;
}

.lifelines {
  margin-top: 30px;
}

.btn-lifeline {
  background: linear-gradient(145deg, var(--warning-color), #e0a800);
  color: var(--primary-color);
  border: none;
  padding: 10px 20px;
  margin: 0 10px;
  border-radius: 25px;
  font-weight: bold;
  transition: all 0.3s ease;
  cursor: pointer;
}

.btn-lifeline:hover:not(.used):not(:disabled) {
  transform: translateY(-3px);
  box-shadow: 0 5px 15px rgba(255, 193, 7, 0.4);
}

.btn-lifeline.used,
.btn-lifeline:disabled {
  background: #666;
  opacity: 0.5;
  cursor: not-allowed;
}

.progress-indicator {
  max-width: 900px;
  margin: 30px auto 0;
}

.progress {
  height: 30px;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 15px;
  overflow: hidden;
}

.progress-bar {
  transition: width 0.5s ease;
  background: linear-gradient(90deg, var(--success-color), #34D399);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
}

.score-display {
  display: flex;
  justify-content: center;
  gap: 20px;
  flex-wrap: wrap;
}

.score-display .badge {
  padding: 10px 20px;
  font-size: 1rem;
}

.btn-outline-warning {
  background: transparent;
  border: 2px solid var(--warning-color);
  color: var(--warning-color);
}

.btn-outline-warning:hover {
  background: var(--warning-color);
  color: var(--primary-color);
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideIn {
  from { transform: translateY(-20px); opacity: 0; }
  to { transform: translateY(0); opacity: 1; }
}

@keyframes correctPulse {
  0%, 100% {
    transform: scale(1);
    box-shadow: 0 0 20px rgba(16, 185, 129, 0.6);
  }
  50% {
    transform: scale(1.05);
    box-shadow: 0 0 40px rgba(16, 185, 129, 1), 0 0 60px rgba(16, 185, 129, 0.8);
  }
}

@keyframes incorrectShake {
  0%, 100% { transform: translateX(0); }
  10%, 30%, 50%, 70%, 90% { transform: translateX(-5px); }
  20%, 40%, 60%, 80% { transform: translateX(5px); }
}

.answer-btn.correct {
  background: linear-gradient(145deg, #10B981, #0C8F63) !important;
  border-color: #10B981 !important;
  animation: correctPulse 1.5s ease-in-out infinite;
}

.answer-btn.incorrect {
  background: linear-gradient(145deg, #CE1126, #A50D1F) !important;
  border-color: #CE1126 !important;
  animation: incorrectShake 0.5s ease-in-out;
}

@media (max-width: 768px) {
  .question-text {
    font-size: 1.3rem;
  }

  .answer-btn {
    padding: 15px;
    font-size: 1rem;
  }
}
</style>
