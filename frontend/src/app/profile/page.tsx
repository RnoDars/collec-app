'use client';

import { useRouter } from 'next/navigation';
import { useAuthStore } from '@/store/auth-store';
import { useEffect } from 'react';

export default function ProfilePage() {
  const router = useRouter();
  const { user, isAuthenticated, logout, isLoading } = useAuthStore();

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      router.push('/auth/login');
    }
  }, [isAuthenticated, isLoading, router]);

  const handleLogout = async () => {
    await logout();
    router.push('/');
  };

  if (isLoading || !user) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-center">
          <div className="text-lg text-gray-600">Chargement...</div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md mx-auto bg-white rounded-lg shadow-md p-8">
        <h1 className="text-2xl font-bold text-gray-900 mb-6">Mon Profil</h1>

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Email
            </label>
            <div className="mt-1 text-gray-900">{user.email}</div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">
              Membre depuis
            </label>
            <div className="mt-1 text-gray-900">
              {new Date(user.createdAt).toLocaleDateString('fr-FR')}
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">
              ID
            </label>
            <div className="mt-1 text-gray-500 text-sm font-mono">
              {user.id}
            </div>
          </div>
        </div>

        <div className="mt-8 space-y-4">
          <button
            onClick={handleLogout}
            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          >
            Se déconnecter
          </button>

          <button
            onClick={() => router.push('/')}
            className="w-full flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Retour à l'accueil
          </button>
        </div>
      </div>
    </div>
  );
}
