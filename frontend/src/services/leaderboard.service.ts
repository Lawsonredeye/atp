import api from './api';
import type { LeaderboardResponse, UserRank, ApiResponse } from '../types';

export interface LeaderboardFilters {
  limit?: number;
  offset?: number;
  period?: 'all_time' | 'weekly' | 'monthly';
  subject_id?: number;
}

export const leaderboardService = {
  async getLeaderboard(filters: LeaderboardFilters = {}): Promise<LeaderboardResponse> {
    const { limit = 50, offset = 0, period = 'all_time', subject_id } = filters;
    let url = `/api/v1/leaderboard?limit=${limit}&offset=${offset}&period=${period}`;
    if (subject_id) {
      url += `&subject_id=${subject_id}`;
    }
    const response = await api.get<ApiResponse<LeaderboardResponse>>(url);
    return response.data.data || { entries: [], period: 'all_time', total_users: 0 };
  },

  async getMyRank(subjectId?: number): Promise<UserRank> {
    let url = '/api/v1/leaderboard/me';
    if (subjectId) {
      url += `?subject_id=${subjectId}`;
    }
    const response = await api.get<ApiResponse<UserRank>>(url);
    return response.data.data!;
  },
};

export default leaderboardService;

