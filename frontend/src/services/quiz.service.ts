import api from './api';
import type { QuizRequest, GeneratedQuiz, SubmitQuizRequest, QuizSubmitResponse, ApiResponse } from '../types';

export const quizService = {
  async createQuiz(request: QuizRequest): Promise<GeneratedQuiz> {
    const response = await api.post<ApiResponse<GeneratedQuiz>>('/api/v1/quiz/create', request);
    return response.data.data!;
  },

  async submitQuiz(answers: SubmitQuizRequest[]): Promise<QuizSubmitResponse> {
    const response = await api.post<ApiResponse<QuizSubmitResponse>>('/api/v1/quiz/submit', answers);
    return response.data.data!;
  },
};

export default quizService;

