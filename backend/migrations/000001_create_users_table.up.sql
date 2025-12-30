-- Migration : Création de la table users
-- Version : 0.2.0
-- Date : 2025-12-30

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index unique sur l'email pour garantir l'unicité
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Commentaires pour documentation
COMMENT ON TABLE users IS 'Table des utilisateurs de l''application';
COMMENT ON COLUMN users.id IS 'Identifiant unique UUID';
COMMENT ON COLUMN users.email IS 'Email unique de l''utilisateur';
COMMENT ON COLUMN users.password IS 'Hash bcrypt du mot de passe';
COMMENT ON COLUMN users.created_at IS 'Date de création du compte';
COMMENT ON COLUMN users.updated_at IS 'Date de dernière modification';
