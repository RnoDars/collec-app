import { authApi, AuthApiError } from '../auth';
import { LoginRequest, RegisterRequest } from '@/types/auth';

// Mock global fetch
global.fetch = jest.fn();

describe('authApi', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('register', () => {
    it('devrait inscrire un utilisateur avec succès', async () => {
      const mockResponse = {
        accessToken: 'mock-access-token',
        refreshToken: 'mock-refresh-token',
        user: {
          id: '123',
          email: 'test@example.com',
          createdAt: '2025-12-30T00:00:00Z',
        },
      };

      (global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => mockResponse,
      });

      const data: RegisterRequest = {
        email: 'test@example.com',
        password: 'password123',
      };

      const result = await authApi.register(data);

      expect(result).toEqual(mockResponse);
      expect(global.fetch).toHaveBeenCalledWith(
        'http://localhost:8080/api/auth/register',
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(data),
        }
      );
    });

    it('devrait lancer une erreur si l\'email existe déjà', async () => {
      const mockError = {
        error: {
          code: 'ERR_AUTH_002',
          message: 'Cet email est déjà utilisé',
        },
      };

      (global.fetch as jest.Mock).mockResolvedValue({
        ok: false,
        status: 409,
        json: async () => mockError,
      });

      const data: RegisterRequest = {
        email: 'existing@example.com',
        password: 'password123',
      };

      await expect(authApi.register(data)).rejects.toThrow(AuthApiError);
      await expect(authApi.register(data)).rejects.toThrow(
        'Cet email est déjà utilisé'
      );
    });
  });

  describe('login', () => {
    it('devrait connecter un utilisateur avec succès', async () => {
      const mockResponse = {
        accessToken: 'mock-access-token',
        refreshToken: 'mock-refresh-token',
        user: {
          id: '123',
          email: 'test@example.com',
          createdAt: '2025-12-30T00:00:00Z',
        },
      };

      (global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => mockResponse,
      });

      const data: LoginRequest = {
        email: 'test@example.com',
        password: 'password123',
      };

      const result = await authApi.login(data);

      expect(result).toEqual(mockResponse);
      expect(global.fetch).toHaveBeenCalledWith(
        'http://localhost:8080/api/auth/login',
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(data),
        }
      );
    });

    it('devrait lancer une erreur si les credentials sont invalides', async () => {
      const mockError = {
        error: {
          code: 'ERR_AUTH_001',
          message: 'Email ou mot de passe incorrect',
        },
      };

      (global.fetch as jest.Mock).mockResolvedValue({
        ok: false,
        status: 401,
        json: async () => mockError,
      });

      const data: LoginRequest = {
        email: 'test@example.com',
        password: 'wrongpassword',
      };

      await expect(authApi.login(data)).rejects.toThrow(AuthApiError);
      await expect(authApi.login(data)).rejects.toThrow(
        'Email ou mot de passe incorrect'
      );
    });
  });

  describe('logout', () => {
    it('devrait déconnecter l\'utilisateur', async () => {
      (global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => ({ message: 'Déconnexion réussie' }),
      });

      await authApi.logout();

      expect(global.fetch).toHaveBeenCalledWith(
        'http://localhost:8080/api/auth/logout',
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
        }
      );
    });
  });

  describe('refreshToken', () => {
    it('devrait refresh le token avec succès', async () => {
      const mockResponse = {
        accessToken: 'new-access-token',
      };

      (global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => mockResponse,
      });

      const result = await authApi.refreshToken({
        refreshToken: 'old-refresh-token',
      });

      expect(result).toEqual(mockResponse);
    });
  });

  describe('getMe', () => {
    it('devrait récupérer les infos utilisateur', async () => {
      const mockUser = {
        id: '123',
        email: 'test@example.com',
        createdAt: '2025-12-30T00:00:00Z',
      };

      (global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => mockUser,
      });

      const result = await authApi.getMe('mock-access-token');

      expect(result).toEqual(mockUser);
      expect(global.fetch).toHaveBeenCalledWith(
        'http://localhost:8080/api/auth/me',
        {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Authorization: 'Bearer mock-access-token',
          },
        }
      );
    });
  });
});
