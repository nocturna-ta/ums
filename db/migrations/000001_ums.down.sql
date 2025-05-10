-- 1. Drop indexes on dependent tables
DROP INDEX IF EXISTS idx_pending_registrations_user_id;
DROP INDEX IF EXISTS idx_users_verification_status;

-- 2. Drop dependent tables in dependency order
DROP TABLE IF EXISTS pending_registrations;
DROP TABLE IF EXISTS voters;
DROP TABLE IF EXISTS kpu_kota;
DROP TABLE IF EXISTS kpu_provinsi;

-- 3. Now it's safe to drop users
DROP TABLE IF EXISTS users;

-- 4. Finally drop the enum types
DROP TYPE IF EXISTS verification_status;
DROP TYPE IF EXISTS users_role;
