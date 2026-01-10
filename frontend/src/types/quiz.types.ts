export interface QuizRequest {
  subject_id: number;
  num_of_questions: number;
}

export interface QuizOption {
  id: number;
  option: string;
}

export interface QuizQuestion {
  question_id: number;
  question: string;
  subject_id: number;
  is_multiple_choice: boolean;
  options: QuizOption[];
}

export interface GeneratedQuiz {
  subject_id: number;
  total_count: number;
  questions: QuizQuestion[];
}

export interface SubmitQuizRequest {
  question_id: number;
  is_multiple_choice: boolean;
  option_ids: number[];
}

export interface QuizResult {
  question_id: number;
  question: string;
  selected_options: string[];
  correct_answer: string;
  is_correct: boolean;
  explanation: string;
}

export interface QuizSubmitResponse {
  user_id: number;
  subject_id: number;
  total_questions: number;
  correct_answers: number;
  incorrect_answers: number;
  score: number;
  results: QuizResult[];
}

