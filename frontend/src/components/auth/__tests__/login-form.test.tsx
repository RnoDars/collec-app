import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { LoginForm } from '../login-form';
import { useAuthStore } from '@/store/auth-store';
import { useRouter } from 'next/navigation';

// Mock Next.js router
jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}));

// Mock auth store
jest.mock('@/store/auth-store');

describe('LoginForm', () => {
  const mockPush = jest.fn();
  const mockLogin = jest.fn();
  const mockClearError = jest.fn();

  beforeEach(() => {
    jest.clearAllMocks();
    (useRouter as jest.Mock).mockReturnValue({
      push: mockPush,
    });
    (useAuthStore as unknown as jest.Mock).mockReturnValue({
      login: mockLogin,
      isLoading: false,
      error: null,
      clearError: mockClearError,
    });
  });

  it('devrait afficher le formulaire de connexion', () => {
    render(<LoginForm />);

    expect(screen.getByText('Connexion')).toBeInTheDocument();
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/mot de passe/i)).toBeInTheDocument();
    expect(
      screen.getByRole('button', { name: /se connecter/i })
    ).toBeInTheDocument();
  });

  it('devrait afficher un lien vers la page d\'inscription', () => {
    render(<LoginForm />);

    const registerLink = screen.getByText(/créez un nouveau compte/i);
    expect(registerLink).toBeInTheDocument();
  });

  it('devrait afficher des erreurs de validation si les champs sont vides', async () => {
    render(<LoginForm />);

    const submitButton = screen.getByRole('button', { name: /se connecter/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/l'email est requis/i)).toBeInTheDocument();
      expect(
        screen.getByText(/le mot de passe est requis/i)
      ).toBeInTheDocument();
    });

    expect(mockLogin).not.toHaveBeenCalled();
  });

  it('devrait afficher une erreur si l\'email est invalide', async () => {
    render(<LoginForm />);

    const emailInput = screen.getByLabelText(/email/i) as HTMLInputElement;
    const passwordInput = screen.getByLabelText(/mot de passe/i);
    const form = emailInput.closest('form')!;

    fireEvent.change(emailInput, { target: { value: 'invalid-email' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });

    // Soumettre le formulaire directement au lieu de cliquer sur le bouton
    fireEvent.submit(form);

    await waitFor(() => {
      expect(screen.getByText(/email invalide/i)).toBeInTheDocument();
    });

    expect(mockLogin).not.toHaveBeenCalled();
  });

  it('devrait appeler login avec les bonnes données et rediriger', async () => {
    mockLogin.mockResolvedValueOnce(undefined);

    render(<LoginForm />);

    const emailInput = screen.getByLabelText(/email/i);
    const passwordInput = screen.getByLabelText(/mot de passe/i);

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });

    const submitButton = screen.getByRole('button', { name: /se connecter/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(mockClearError).toHaveBeenCalled();
      expect(mockLogin).toHaveBeenCalledWith({
        email: 'test@example.com',
        password: 'password123',
      });
      expect(mockPush).toHaveBeenCalledWith('/');
    });
  });

  it('devrait afficher un message d\'erreur si la connexion échoue', async () => {
    mockLogin.mockRejectedValueOnce(new Error('Login failed'));

    render(<LoginForm />);

    const emailInput = screen.getByLabelText(/email/i);
    const passwordInput = screen.getByLabelText(/mot de passe/i);

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(passwordInput, { target: { value: 'wrongpassword' } });

    const submitButton = screen.getByRole('button', { name: /se connecter/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(
        screen.getByText(/connexion échouée/i)
      ).toBeInTheDocument();
    });
  });

  it('devrait désactiver le bouton pendant le chargement', () => {
    (useAuthStore as unknown as jest.Mock).mockReturnValue({
      login: mockLogin,
      isLoading: true,
      error: null,
      clearError: mockClearError,
    });

    render(<LoginForm />);

    const submitButton = screen.getByRole('button', { name: /connexion/i });
    expect(submitButton).toBeDisabled();
    expect(submitButton).toHaveTextContent('Connexion...');
  });

  it('devrait afficher une erreur du store si présente', () => {
    (useAuthStore as unknown as jest.Mock).mockReturnValue({
      login: mockLogin,
      isLoading: false,
      error: 'Email ou mot de passe incorrect',
      clearError: mockClearError,
    });

    render(<LoginForm />);

    expect(
      screen.getByText(/email ou mot de passe incorrect/i)
    ).toBeInTheDocument();
  });
});
