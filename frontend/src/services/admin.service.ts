import api from './api';
import type {
  ApiResponse,
  AuthTokens,
  AdminLoginRequest,
  QuestionData,
  CreateSubjectRequest,
  Subject,
  QuestionWithDetails
} from '../types';

export const adminService = {
  // Auth
  async login(credentials: AdminLoginRequest): Promise<AuthTokens> {
    const response = await api.post<ApiResponse<AuthTokens>>('/admin/login', credentials);
    const tokens = response.data.data!;
    localStorage.setItem('access_token', tokens.access_token);
    localStorage.setItem('refresh_token', tokens.refresh_token);
    localStorage.setItem('is_admin', 'true');
    return tokens;
  },

  logout(): void {
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    localStorage.removeItem('is_admin');
  },

  isAdmin(): boolean {
    return localStorage.getItem('is_admin') === 'true';
  },

  // Questions
  async createBulkQuestions(subjectId: number, questions: QuestionData[]): Promise<void> {
    await api.post(`/api/v1/admin/questions/bulk/${subjectId}`, questions);
  },

  async createSingleQuestion(subjectId: number, question: QuestionData): Promise<void> {
    await api.post(`/api/v1/admin/questions/single/${subjectId}`, question);
  },

  async getAllQuestions(): Promise<QuestionWithDetails[]> {
    const response = await api.get<ApiResponse<QuestionWithDetails[]>>('/api/v1/admin/questions');
    return response.data.data || [];
  },

  async getQuestionById(id: number): Promise<QuestionWithDetails> {
    const response = await api.get<ApiResponse<QuestionWithDetails>>(`/api/v1/admin/questions/${id}`);
    return response.data.data!;
  },

  async deleteQuestion(id: number): Promise<void> {
    await api.delete(`/api/v1/admin/questions/${id}`);
  },

  // Subjects
  async getAllSubjects(): Promise<Subject[]> {
    const response = await api.get<ApiResponse<Subject[]>>('/api/v1/admin/subject');
    return response.data.data || [];
  },

  async createSubject(data: CreateSubjectRequest): Promise<Subject> {
    const response = await api.post<ApiResponse<Subject>>('/api/v1/admin/subject', data);
    return response.data.data!;
  },
};

export default adminService;

