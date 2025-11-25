import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '@/services/api'
import type { Question, GameResults, AnswerSubmission } from '@/types'

export const useGameStore = defineStore('game', () => {
  const sessionId = ref<string | null>(null)
  const currentQuestion = ref<Question | null>(null)
  const questionNumber = ref<number>(0)
  const totalQuestions = ref<number>(0)
  const timeRemaining = ref<number>(0)
  const startTime = ref<number | null>(null)

  // Game stats
  const correctAnswers = ref<number>(0)
  const incorrectAnswers = ref<number>(0)
  const flaggedCount = ref<number>(0)
  const flaggedQuestions = ref<Set<number>>(new Set())

  // Lifelines
  const fiftyFiftyUsed = ref<boolean>(false)
  const hintUsed = ref<boolean>(false)
  const skipsRemaining = ref<number>(3)

  // Game config
  const gameMode = ref<string | null>(null)
  const categories = ref<string[]>([])
  const timeLimit = ref<number>(0)

  // Results
  const results = ref<GameResults | null>(null)

  async function startGame(
    mode: string,
    questionCount: number,
    timeLimitMinutes: number,
    selectedCategories: string[],
    difficulty?: string,
    focusWeakAreas?: boolean
  ): Promise<{ success: boolean; message?: string }> {
    try {
      const data = await api.startGame(
        mode,
        questionCount,
        timeLimitMinutes,
        selectedCategories,
        difficulty,
        focusWeakAreas
      )

      sessionId.value = data.session_id
      gameMode.value = mode
      categories.value = data.config.categories
      timeLimit.value = timeLimitMinutes
      totalQuestions.value = data.config.question_count
      startTime.value = Date.now()

      // Reset counters
      correctAnswers.value = 0
      incorrectAnswers.value = 0
      flaggedCount.value = 0
      flaggedQuestions.value.clear()
      fiftyFiftyUsed.value = false
      hintUsed.value = false
      skipsRemaining.value = 3

      return { success: true }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al iniciar el juego' }
    }
  }

  async function loadNextQuestion(): Promise<{ success: boolean; noMoreQuestions?: boolean; message?: string }> {
    try {
      const data = await api.getQuestion(sessionId.value!)
      currentQuestion.value = data.question
      questionNumber.value = data.question_number
      totalQuestions.value = data.total_questions
      timeRemaining.value = data.time_remaining
      return { success: true }
    } catch (error: any) {
      if (error.response?.status === 404) {
        return { success: false, noMoreQuestions: true }
      }
      return { success: false, message: error.response?.data?.error || 'Error al cargar la pregunta' }
    }
  }

  async function submitAnswer(
    choiceId: number
  ): Promise<{ success: boolean; correct?: boolean; correctChoiceId?: number; explanation?: string; message?: string }> {
    try {
      const timeSpent = Math.floor((Date.now() - startTime.value!) / 1000)
      const submission: AnswerSubmission = {
        question_id: currentQuestion.value!.id,
        choice_id: choiceId,
        time_spent: timeSpent
      }

      const data = await api.submitAnswer(sessionId.value!, submission)

      if (data.correct) {
        correctAnswers.value++
      } else {
        incorrectAnswers.value++
      }

      return {
        success: true,
        correct: data.correct,
        correctChoiceId: data.correct_choice_id,
        explanation: data.explanation
      }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al enviar la respuesta' }
    }
  }

  async function flagQuestion(): Promise<{ success: boolean; isFlagged?: boolean; message?: string }> {
    try {
      const questionId = currentQuestion.value!.id
      await api.flagQuestion(sessionId.value!, { question_id: questionId })

      // Toggle the flag state
      if (flaggedQuestions.value.has(questionId)) {
        flaggedQuestions.value.delete(questionId)
        flaggedCount.value--
      } else {
        flaggedQuestions.value.add(questionId)
        flaggedCount.value++
      }

      return { success: true, isFlagged: flaggedQuestions.value.has(questionId) }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al marcar la pregunta' }
    }
  }

  async function endGame(): Promise<{ success: boolean; message?: string }> {
    try {
      await api.endGame(sessionId.value!)
      return { success: true }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al finalizar el juego' }
    }
  }

  async function loadResults(): Promise<{ success: boolean; data?: GameResults; message?: string }> {
    try {
      const data = await api.getResults(sessionId.value!)
      results.value = data
      return { success: true, data }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al cargar los resultados' }
    }
  }

  function useFiftyFifty(): boolean {
    if (!fiftyFiftyUsed.value && currentQuestion.value) {
      fiftyFiftyUsed.value = true
      return true
    }
    return false
  }

  function useHint(): boolean {
    if (!hintUsed.value && currentQuestion.value) {
      hintUsed.value = true
      return true
    }
    return false
  }

  function useSkip(): boolean {
    if (skipsRemaining.value > 0) {
      skipsRemaining.value--
      return true
    }
    return false
  }

  function resetGame(): void {
    sessionId.value = null
    currentQuestion.value = null
    questionNumber.value = 0
    totalQuestions.value = 0
    timeRemaining.value = 0
    startTime.value = null
    correctAnswers.value = 0
    incorrectAnswers.value = 0
    flaggedCount.value = 0
    flaggedQuestions.value.clear()
    fiftyFiftyUsed.value = false
    hintUsed.value = false
    skipsRemaining.value = 3
    gameMode.value = null
    categories.value = []
    timeLimit.value = 0
    results.value = null
  }

  function isQuestionFlagged(questionId: number): boolean {
    return flaggedQuestions.value.has(questionId)
  }

  return {
    sessionId,
    currentQuestion,
    questionNumber,
    totalQuestions,
    timeRemaining,
    startTime,
    correctAnswers,
    incorrectAnswers,
    flaggedCount,
    flaggedQuestions,
    fiftyFiftyUsed,
    hintUsed,
    skipsRemaining,
    gameMode,
    categories,
    timeLimit,
    results,
    startGame,
    loadNextQuestion,
    submitAnswer,
    flagQuestion,
    endGame,
    loadResults,
    useFiftyFifty,
    useHint,
    useSkip,
    resetGame,
    isQuestionFlagged
  }
})
