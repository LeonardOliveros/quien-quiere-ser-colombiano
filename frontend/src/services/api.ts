import axios, { type AxiosInstance, type AxiosError } from 'axios'
import type {
  LoginCredentials,
  RegisterCredentials,
  AuthResponse,
  RegisterResponse,
  StartGameResponse,
  QuestionResponse,
  AnswerSubmission,
  AnswerResponse,
  FlagRequest,
  GameResults,
  UserStats,
  WeakArea,
  GameHistory,
  Recommendation,
  Question,
  QuestionCount
} from '@/types'

class ApiService {
  private axiosInstance: AxiosInstance

  constructor() {
    this.axiosInstance = axios.create({
      baseURL: '/api',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // Request interceptor to add auth token
    this.axiosInstance.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('token')
        if (token) {
          // Backend expects token without "Bearer " prefix
          config.headers.Authorization = token
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      }
    )

    // Response interceptor to handle errors
    this.axiosInstance.interceptors.response.use(
      (response) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401) {
          localStorage.removeItem('token')
          localStorage.removeItem('userId')
          window.location.href = '/login'
        }
        return Promise.reject(error)
      }
    )
  }

  // Auth endpoints
  async register(credentials: RegisterCredentials): Promise<RegisterResponse> {
    const response = await this.axiosInstance.post<RegisterResponse>('/register', credentials)
    return response.data
  }

  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    const response = await this.axiosInstance.post<AuthResponse>('/login', credentials)
    return response.data
  }

  // Game endpoints
  async startGame(
    mode: string,
    questionCount: number,
    timeLimit: number,
    categories: string[],
    difficulty?: string,
    focusWeakAreas?: boolean
  ): Promise<StartGameResponse> {
    const response = await this.axiosInstance.post<StartGameResponse>('/game/start', {
      mode,
      question_count: questionCount,
      time_limit: timeLimit,
      categories,
      difficulty,
      focus_weak_areas: focusWeakAreas
    })
    return response.data
  }

  async getQuestion(sessionId: string): Promise<QuestionResponse> {
    const response = await this.axiosInstance.get<QuestionResponse>(`/game/${sessionId}/question`)
    return response.data
  }

  async submitAnswer(sessionId: string, submission: AnswerSubmission): Promise<AnswerResponse> {
    const response = await this.axiosInstance.post<AnswerResponse>(
      `/game/${sessionId}/answer`,
      submission
    )
    return response.data
  }

  async flagQuestion(sessionId: string, request: FlagRequest): Promise<{ message: string }> {
    const response = await this.axiosInstance.post<{ message: string }>(
      `/game/${sessionId}/flag`,
      request
    )
    return response.data
  }

  async useFiftyFifty(sessionId: string, questionId: number): Promise<{ remove_choice_ids: number[] }> {
    const response = await this.axiosInstance.post<{ remove_choice_ids: number[] }>(
      `/game/${sessionId}/fifty-fifty`,
      { question_id: questionId }
    )
    return response.data
  }

  async endGame(sessionId: string): Promise<{ message: string }> {
    const response = await this.axiosInstance.post<{ message: string }>(`/game/${sessionId}/end`)
    return response.data
  }

  async pauseGame(sessionId: string): Promise<{ message: string }> {
    const response = await this.axiosInstance.post<{ message: string }>(`/game/${sessionId}/pause`)
    return response.data
  }

  async getPausedGame(mode: string): Promise<any> {
    const response = await this.axiosInstance.get(`/game/paused/${mode}`)
    return response.data
  }

  async getResults(sessionId: string): Promise<GameResults> {
    const response = await this.axiosInstance.get<GameResults>(`/game/${sessionId}/results`)
    return response.data
  }

  // User statistics endpoints
  async getUserStats(userId: string): Promise<UserStats> {
    const response = await this.axiosInstance.get<UserStats>(`/user/${userId}/stats`)
    return response.data
  }

  async getWeakAreas(userId: string): Promise<{ weak_areas: WeakArea[] }> {
    const response = await this.axiosInstance.get<{ weak_areas: WeakArea[] }>(
      `/user/${userId}/weak-areas`
    )
    return response.data
  }

  async getUserHistory(userId: string): Promise<GameHistory[]> {
    const response = await this.axiosInstance.get<GameHistory[]>(`/user/${userId}/history`)
    return response.data
  }

  async resetStats(userId: string): Promise<{ message: string }> {
    const response = await this.axiosInstance.delete<{ message: string }>(`/user/${userId}/stats`)
    return response.data
  }

  // Recommendations endpoint
  async getRecommendations(userId: string): Promise<Recommendation[]> {
    const response = await this.axiosInstance.get<Recommendation[]>(`/recommendations/${userId}`)
    return response.data
  }

  // Questions endpoints
  async getAllQuestions(): Promise<Question[]> {
    const response = await this.axiosInstance.get<Question[]>('/questions')
    return response.data
  }

  async getQuestionById(id: number): Promise<Question> {
    const response = await this.axiosInstance.get<Question>(`/questions/${id}`)
    return response.data
  }

  async getQuestionsByCategory(category: string): Promise<Question[]> {
    const response = await this.axiosInstance.get<Question[]>(`/questions/category/${category}`)
    return response.data
  }

  async getQuestionCount(): Promise<QuestionCount> {
    const response = await this.axiosInstance.get<QuestionCount>('/questions/count')
    return response.data
  }
}

export default new ApiService()
