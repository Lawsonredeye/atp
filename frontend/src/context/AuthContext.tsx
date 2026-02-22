import { createContext, useContext, useReducer, useEffect, type ReactNode } from 'react';
import { authService } from '../services';
import type { AuthState, AuthTokens, LoginRequest, RegisterRequest, UserDashboard, User } from '../types';

type AuthAction =
  | { type: 'AUTH_START' }
  | { type: 'AUTH_SUCCESS'; payload: { user: UserDashboard; tokens: AuthTokens } }
  | { type: 'AUTH_FAILURE' }
  | { type: 'LOGOUT' }
  | { type: 'SET_USER'; payload: UserDashboard }
  | { type: 'REGISTER_SUCCESS' };

const initialState: AuthState = {
  isAuthenticated: false,
  user: null,
  tokens: null,
  loading: true,
};

function authReducer(state: AuthState, action: AuthAction): AuthState {
  switch (action.type) {
    case 'AUTH_START':
      return { ...state, loading: true };
    case 'AUTH_SUCCESS':
      return { ...state, isAuthenticated: true, user: action.payload.user, tokens: action.payload.tokens, loading: false };
    case 'AUTH_FAILURE':
      return { ...state, isAuthenticated: false, user: null, tokens: null, loading: false };
    case 'LOGOUT':
      return { ...initialState, loading: false };
    case 'SET_USER':
      return { ...state, user: action.payload, loading: false };
    case 'REGISTER_SUCCESS':
      return { ...state, loading: false };
    default:
      return state;
  }
}

interface AuthContextType {
  state: AuthState;
  login: (credentials: LoginRequest) => Promise<void>;
  register: (userData: RegisterRequest) => Promise<User>;
  logout: () => void;
  refreshUser: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(authReducer, initialState);

  useEffect(() => {
    const initAuth = async () => {
      const accessToken = authService.getAccessToken();
      const refreshToken = authService.getRefreshToken();
      if (accessToken && refreshToken) {
        try {
          const dashboard = await authService.getDashboard();
          dispatch({ type: 'AUTH_SUCCESS', payload: { user: dashboard, tokens: { access_token: accessToken, refresh_token: refreshToken } } });
        } catch {
          authService.logout();
          dispatch({ type: 'AUTH_FAILURE' });
        }
      } else {
        dispatch({ type: 'AUTH_FAILURE' });
      }
    };
    initAuth();
  }, []);

  const login = async (credentials: LoginRequest) => {
    dispatch({ type: 'AUTH_START' });
    try {
      const tokens = await authService.login(credentials);
      const dashboard = await authService.getDashboard();
      dispatch({ type: 'AUTH_SUCCESS', payload: { user: dashboard, tokens } });
    } catch (error) {
      dispatch({ type: 'AUTH_FAILURE' });
      throw error;
    }
  };

  const register = async (userData: RegisterRequest): Promise<User> => {
    dispatch({ type: 'AUTH_START' });
    try {
      const user = await authService.register(userData);
      dispatch({ type: 'REGISTER_SUCCESS' });
      return user;
    } catch (error) {
      dispatch({ type: 'AUTH_FAILURE' });
      throw error;
    }
  };

  const logout = () => {
    authService.logout();
    dispatch({ type: 'LOGOUT' });
  };

  const refreshUser = async () => {
    try {
      const dashboard = await authService.getDashboard();
      dispatch({ type: 'SET_USER', payload: dashboard });
    } catch (error) {
      console.error('Failed to refresh user:', error);
    }
  };

  return (
    <AuthContext.Provider value={{ state, login, register, logout, refreshUser }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
