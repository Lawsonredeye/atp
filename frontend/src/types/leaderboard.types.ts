export interface LeaderboardEntry {
  rank: number;
  user_id: number;
  user_name: string;
  total_score: number;
  total_quizzes: number;
  correct_answers: number;
  total_questions: number;
  accuracy_percent: number;
}

export interface LeaderboardResponse {
  subject_id?: number;
  subject_name?: string;
  period: string;
  total_users: number;
  entries: LeaderboardEntry[];
}

export interface UserRank {
  user_id: number;
  user_name: string;
  rank: number;
  total_score: number;
  total_quizzes: number;
  correct_answers: number;
  total_questions: number;
  accuracy_percent: number;
  total_users: number;
}

