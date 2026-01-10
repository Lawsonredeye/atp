// Admin-specific types

export interface AdminLoginRequest {
  email: string;
  password: string;
}

export interface QuestionData {
  name: string;
  options: string[];
  answer: string;
  explanation: string;
}

export interface CreateSubjectRequest {
  name: string;
}

export interface QuestionWithDetails {
  id: number;
  subject_id: number;
  name: string;
  is_multiple_choice: boolean;
  created_at: string;
  updated_at: string;
  options: {
    id: number;
    option: string;
  }[];
  answer?: {
    id: number;
    answer: string;
    explanation: string;
  };
}

