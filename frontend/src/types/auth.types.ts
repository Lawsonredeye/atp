export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  full_name: string;
  email: string;
  password: string;
}

export interface AuthTokens {
  access_token: string;
  refresh_token: string;
}

export interface AuthState {
  isAuthenticated: boolean;
  user: UserDashboard | null;
  tokens: AuthTokens | null;
  loading: boolean;
}

export interface User {
  id: number;
  full_name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface UserDashboard {
  id: number;
  full_name: string;
  email: string;
  created_at: string;
  updated_at: string;
  user_id: number;
  total_quizzes_taken: number;
  total_questions_answered: number;
  total_correct_answers: number;
  total_incorrect_answers: number;
  average_accuracy: number;
  roles: string[];
}

