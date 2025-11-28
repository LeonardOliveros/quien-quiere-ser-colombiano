// Auth types
export interface LoginCredentials {
  username: string
  password: string
}

export interface RegisterCredentials {
  username: string
  email: string
  password: string
}

export interface AuthResponse {
  user_id: string
  token: string
  message: string
}

export interface RegisterResponse {
  user_id: string
  message: string
}

// Question types
export interface Choice {
  id: number
  text: string
}

export interface Question {
  id: number
  text: string
  category: string
  subcategory: string
  difficulty: string
  hint?: string
  explanation?: string
  choices: Choice[]
}

export interface QuestionResponse {
  question: Question
  question_number: number
  total_questions: number
  time_remaining: number
  time_elapsed: number
}

// Game types
export interface GameConfig {
  mode: string
  question_count: number
  time_limit: number
  categories: string[]
  difficulty?: string
  focus_weak_areas?: boolean
}

export interface StartGameResponse {
  session_id: string
  config: GameConfig
  message: string
}

export interface AnswerSubmission {
  question_id: number
  choice_id: number
  time_spent: number
}

export interface AnswerResponse {
  correct: boolean
  choice_id: number
  correct_choice_id: number
  explanation: string
}

export interface FlagRequest {
  question_id: number
}

// Results types
export interface CategoryScore {
  category: string
  total_questions: number
  correct_answers: number
  percentage: number
  passed: boolean
}

export interface IncorrectAnswer {
  question: Question
  user_choice: Choice
  correct_choice: Choice
  explanation: string
}

export interface FlaggedQuestion {
  id: number
  text: string
  category: string
  subcategory: string
  difficulty: string
  hint?: string
  explanation?: string
  choices: Choice[]
}

export interface Recommendation {
  category: string
  subcategory: string
  description: string
  priority: number
  resources?: string
}

export interface GameResults {
  session_id: string
  total_questions: number
  correct_answers: number
  score: number
  percentage: number
  time_taken: number
  category_scores: { [key: string]: CategoryScore }
  incorrect_answers: IncorrectAnswer[]
  flagged_questions: FlaggedQuestion[]
  recommendations: string[]
}

// Statistics types
export interface WeakArea {
  category: string
  subcategory: string
  accuracy: number
  total_attempts: number
}

export interface StrongArea {
  category: string
  subcategory: string
  accuracy: number
  total_attempts: number
}

export interface RecentProgress {
  date: string
  score: number
}

export interface CategoryStat {
  category: string
  total_questions: number
  correct_answers: number
  average_percentage: number
  improvement?: number
}

export interface CategoryStats {
  [key: string]: CategoryStat
}

export interface UserStats {
  total_games: number
  average_score: number
  best_score: number
  total_questions: number
  category_stats: CategoryStats
  weak_areas: WeakArea[]
  strong_areas: StrongArea[]
  recent_progress: RecentProgress[]
}

export interface GameHistory {
  created_at: string
  mode: string
  correct_answers: number
  total_questions: number
  score: number
}

export interface QuestionCount {
  total: number
  by_category: {
    [key: string]: number
  }
  by_subcategory: Array<{
    category: string
    subcategory: string
    count: number
  }>
}

// API Response types
export interface ApiResponse<T> {
  success: boolean
  data?: T
  message?: string
  error?: string
}
