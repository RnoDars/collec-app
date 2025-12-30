import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { RegisterForm } from '../register-form';
import { useAuthStore } from '@/store/auth-store';
import { useRouter } from 'next/navigation';

// Mock Next.js router
jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}));

// Mock auth store
jest.mock('@/store/auth-store');

describe('RegisterForm', () => {
  const mockPush = jest.fn();
  const mockRegister = jest.fn();
  const mockClearError = jest.fn();

  beforeEach(() => {
    jest.clearAllMocks();
    (useRouter as jest.Mock).mockReturnValue({
      push: mockPush,
    });
    (useAuthStore as unknown as jest.Mock).mockReturnValue({
      register: mockRegister,
      isLoading: false,
      error: null,
      clearError: mockClearError,
    });
  });

  it('devrait afficher le formulaire d\'inscription', () => {
    render(<RegisterForm />);

    expect(screen.getByText('Créer un compte')).toBeInTheDocument();
    expect(screen.getByLabelText(/^email$/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/^mot de passe$/i)).toBeInTheDocument();
    expect(
      screen.getByLabelText(/confirmer le mot de passe/i)
    ).toBeInTheDocument();
    expect(
      screen.getByRole('button', { name: /s'inscrire/i })
    ).toBeInTheDocument();
  });

  it('devrait afficher un lien vers la page de connexion', () => {
    render(<RegisterForm />);

    const loginLink = screen.getByText(/connectez-vous à votre compte existant/i);
    expect(loginLink).toBeInTheDocument();
  });

  it('devrait afficher des erreurs de validation si les champs sont vides', async () => {
    render(<RegisterForm />);

    const form = screen.getByRole('button', { name: /s'inscrire/i }).closest('form')!;
    fireEvent.submit(form);

    await waitFor(() => {
      expect(screen.getByText(/l'email est requis/i)).toBeInTheDocument();
      expect(
        screen.getByText(/le mot de passe doit contenir au moins 8 caractères/i)
      ).toBeInTheDocument();
    });

    expect(mockRegister).not.toHaveBeenCalled();
  });

  it('devrait afficher une erreur si l\'email est invalide', async () => {
    render(<RegisterForm />);

    const emailInput = screen.getByLabelText(/^email$/i) as HTMLInputElement;
    const passwordInput = screen.getByLabelText(/^mot de passe$/i);
    const confirmPasswordInput = screen.getByLabelText(/confirmer le mot de passe/i);
    const form = emailInput.closest('form')!;

    fireEvent.change(emailInput, { target: { value: 'invalid-email' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });
    fireEvent.change(confirmPasswordInput, { target: { value: 'password123' } });

    fireEvent.submit(form);

    await waitFor(() => {
      expect(screen.getByText(/email invalide/i)).toBeInTheDocument();
    });

    expect(mockRegister).not.toHaveBeenCalled();
  });

  it('devrait afficher une erreur si le mot de passe est trop court', async () => {
    render(<RegisterForm />);

    const emailInput = screen.getByLabelText(/^email$/i);
    const passwordInput = screen.getByLabelText(/^mot de passe$/i);
    const confirmPasswordInput = screen.getByLabelText(/confirmer le mot de passe/i);
    const form = emailInput.closest('form')!;

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(passwordInput, { target: { value: 'short' } });
    fireEvent.change(confirmPasswordInput, { target: { value: 'short' } });

    fireEvent.submit(form);

    await waitFor(() => {
      expect(
        screen.getByText(/le mot de passe doit contenir au moins 8 caractères/i)
      ).toBeInTheDocument();
    });

    expect(mockRegister).not.toHaveBeenCalled();
  });

  it('devrait afficher une erreur si les mots de passe ne correspondent pas', async () => {
    render(<RegisterForm />);

    const emailInput = screen.getByLabelText(/^email$/i);
    const passwordInput = screen.getByLabelText(/^mot de passe$/i);
    const confirmPasswordInput = screen.getByLabelText(/confirmer le mot de passe/i);
    const form = emailInput.closest('form')!;

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });
    fireEvent.change(confirmPasswordInput, { target: { value: 'different123' } });

    fireEvent.submit(form);

    await waitFor(() => {
      expect(
        screen.getByText(/les mots de passe ne correspondent pas/i)
      ).toBeInTheDocument();
    });

    expect(mockRegister).not.toHaveBeenCalled();
  });

  it('devrait appeler register avec les bonnes données et rediriger', async () => {
    mockRegister.mockResolvedValueOnce(undefined);

    render(<RegisterForm />);

    const emailInput = screen.getByLabelText(/^email$/i);
    const passwordInput = screen.getByLabelText(/^mot de passe$/i);
    const confirmPasswordInput = screen.getByLabelText(/confirmer le mot de passe/i);

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });
    fireEvent.change(confirmPasswordInput, { target: { value: 'password123' } });

    const form = emailInput.closest('form')!;
    fireEvent.submit(form);

    await waitFor(() => {
      expect(mockClearError).toHaveBeenCalled();
      expect(mockRegister).toHaveBeenCalledWith({
        email: 'test@example.com',
        password: 'password123',
      });
      expect(mockPush).toHaveBeenCalledWith('/');
    });
  });

  it('devrait afficher un message d\'erreur si l\'inscription échoue', async () => {
    mockRegister.mockRejectedValueOnce(new Error('Registration failed'));

    render(<RegisterForm />);

    const emailInput = screen.getByLabelText(/^email$/i);
    const passwordInput = screen.getByLabelText(/^mot de passe$/i);
    const confirmPasswordInput = screen.getByLabelText(/confirmer le mot de passe/i);

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });
    fireEvent.change(confirmPasswordInput, { target: { value: 'password123' } });

    const form = emailInput.closest('form')!;
    fireEvent.submit(form);

    await waitFor(() => {
      expect(
        screen.getByText(/inscription échouée/i)
      ).toBeInTheDocument();
    });
  });

  it('devrait désactiver le bouton pendant le chargement', () => {
    (useAuthStore as unknown as jest.Mock).mockReturnValue({
      register: mockRegister,
      isLoading: true,
      error: null,
      clearError: mockClearError,
    });

    render(<RegisterForm />);

    const submitButton = screen.getByRole('button', { name: /inscription/i });
    expect(submitButton).toBeDisabled();
    expect(submitButton).toHaveTextContent('Inscription...');
  });

  it('devrait afficher une erreur du store si présente', () => {
    (useAuthStore as unknown as jest.Mock).mockReturnValue({
      register: mockRegister,
      isLoading: false,
      error: 'Cet email est déjà utilisé',
      clearError: mockClearError,
    });

    render(<RegisterForm />);

    expect(
      screen.getByText(/cet email est déjà utilisé/i)
    ).toBeInTheDocument();
  });
});
