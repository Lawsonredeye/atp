import api from './api';
import type { LeaderboardEntry, LeaderboardResponse, UserRank, ApiResponse } from '../types';

export const leaderboardService = {
  async getLeaderboard(limit = 50, offset = 0): Promise<LeaderboardEntry[]> {
    const response = await api.get<ApiResponse<LeaderboardResponse>>(`/api/v1/leaderboard?limit=${limit}&offset=${offset}`);
    return response.data.data?.entries || [];
  },

  async getMyRank(): Promise<UserRank> {
    const response = await api.get<ApiResponse<UserRank>>('/api/v1/leaderboard/me');
    return response.data.data!;
  },
};

export default leaderboardService;

