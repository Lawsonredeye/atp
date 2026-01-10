import api from './api';
import type { LoginRequest, RegisterRequest, AuthTokens, UserDashboard, ApiResponse } from '../types';

export const authService = {
  async login(credentials: LoginRequest): Promise<AuthTokens> {
    const response = await api.post<ApiResponse<AuthTokens>>('/user/login', credentials);
    const tokens = response.data.data!;
    localStorage.setItem('access_token', tokens.access_token);
    localStorage.setItem('refresh_token', tokens.refresh_token);
    return tokens;
  },

  async register(userData: RegisterRequest): Promise<AuthTokens> {
    const response = await api.post<ApiResponse<AuthTokens>>('/user/register', userData);
    const tokens = response.data.data!;
    localStorage.setItem('access_token', tokens.access_token);
    localStorage.setItem('refresh_token', tokens.refresh_token);
    return tokens;
  },

  async getDashboard(): Promise<UserDashboard> {
    const response = await api.get<ApiResponse<UserDashboard>>('/api/v1/dashboard');
    return response.data.data!;
  },

  logout(): void {
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
  },

  isAuthenticated(): boolean {
    return !!localStorage.getItem('access_token');
  },

  getAccessToken(): string | null {
    return localStorage.getItem('access_token');
  },

  getRefreshToken(): string | null {
    return localStorage.getItem('refresh_token');
  },
};

export default authService;

