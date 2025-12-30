'use client';

import Link from 'next/link';
import { useAuthStore } from '@/store/auth-store';

export default function Home() {
  const { user, isAuthenticated } = useAuthStore();

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="text-center">
        <h1 className="text-5xl font-bold text-gray-900 mb-4">
          Collec-App
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          Application de gestion de collections
        </p>

        <div className="bg-white rounded-lg shadow-lg p-8 max-w-2xl">
          {isAuthenticated && user ? (
            <>
              <h2 className="text-2xl font-semibold text-gray-800 mb-4">
                Bienvenue, {user.email} !
              </h2>
              <p className="text-gray-600 mb-6">
                Vous êtes connecté à votre compte.
              </p>
              <div className="flex gap-4 justify-center">
                <Link
                  href="/profile"
                  className="px-6 py-3 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition"
                >
                  Mon profil
                </Link>
              </div>
            </>
          ) : (
            <>
              <h2 className="text-2xl font-semibold text-gray-800 mb-4">
                Version 0.2.0 - Authentification
              </h2>
              <p className="text-gray-600 mb-6">
                Créez un compte ou connectez-vous pour commencer.
              </p>
              <div className="flex gap-4 justify-center">
                <Link
                  href="/auth/register"
                  className="px-6 py-3 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition"
                >
                  S'inscrire
                </Link>
                <Link
                  href="/auth/login"
                  className="px-6 py-3 bg-white text-blue-600 border-2 border-blue-600 rounded-md hover:bg-blue-50 transition"
                >
                  Se connecter
                </Link>
              </div>
            </>
          )}

          <div className="mt-6 text-left bg-gray-50 p-4 rounded">
            <h3 className="font-semibold text-gray-700 mb-2">Prochaines étapes :</h3>
            <ul className="list-disc list-inside text-gray-600 space-y-1">
              <li>✅ v0.2.0 : Authentification et gestion utilisateurs</li>
              <li>v0.3.0 : Gestion des collections</li>
              <li>v0.4.0 : Gestion des items</li>
              <li>v0.5.0 : Catégories et tags</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
