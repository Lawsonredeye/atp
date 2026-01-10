import api from './api';
import type { Subject, ApiResponse } from '../types';

export const subjectService = {
  async getAllSubjects(): Promise<Subject[]> {
    const response = await api.get<ApiResponse<Subject[]>>('/api/v1/admin/subject');
    return response.data.data || [];
  },

  async getSubjectById(id: number): Promise<Subject> {
    const response = await api.get<ApiResponse<Subject>>(`/api/v1/admin/subject/${id}`);
    return response.data.data!;
  },
};

export default subjectService;

