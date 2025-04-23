-- Drop index first
DROP INDEX IF EXISTS idx_users_verification_status;

-- Drop dependent tables
DROP TABLE IF EXISTS voters;
DROP TABLE IF EXISTS kpu_kota;
DROP TABLE IF EXISTS kpu_provinsi;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop enum types
DROP TYPE IF EXISTS verification_status;
DROP TYPE IF EXISTS users_role;