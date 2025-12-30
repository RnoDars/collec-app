-- Migration rollback : Suppression de la table users
-- Version : 0.2.0
-- Date : 2025-12-30

DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;
