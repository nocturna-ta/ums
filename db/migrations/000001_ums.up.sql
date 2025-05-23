-- Create user role enum
CREATE TYPE users_role AS ENUM('voter', 'unverified', 'kpu_kota', 'kpu_provinsi', 'kpu_pusat');

-- Create verification status enum
CREATE TYPE verification_status AS ENUM('pending', 'approved', 'rejected');

-- Create users table with verification fields
CREATE TABLE "users" (
    id uuid NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    password_salt VARCHAR(10) NOT NULL,
    role users_role NOT NULL,
    is_active BOOL NOT NULL,
    requested_role users_role NULL,
    verification_status verification_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted BOOL NOT NULL DEFAULT FALSE
);

-- Create index on verification status
CREATE INDEX idx_users_verification_status ON users (verification_status);

-- Create KPU Kota table
CREATE TABLE "kpu_kota" (
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE,
    username VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    is_active BOOL NOT NULL,
    photo_path VARCHAR(255) DEFAULT NULL,
    telephone VARCHAR(20) DEFAULT '',
    registered_at TIMESTAMP(6) WITH TIME ZONE DEFAULT now(),
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted BOOL NOT NULL DEFAULT FALSE
);

-- Create KPU Provinsi table
CREATE TABLE "kpu_provinsi" (
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE,
    username VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    is_active BOOL NOT NULL,
    photo_path VARCHAR(255) DEFAULT NULL,
    telephone VARCHAR(20) DEFAULT '',
    registered_at TIMESTAMP(6) WITH TIME ZONE DEFAULT now(),
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted BOOL NOT NULL DEFAULT FALSE
);

-- Create voters table
CREATE TABLE "voters" (
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE,
    nik VARCHAR(16) NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    birth_place VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    residential_address TEXT,
    region VARCHAR(255) NOT NULL,
    voter_address VARCHAR(255) NOT NULL,
    is_registered BOOL NOT NULL DEFAULT FALSE,
    has_voted BOOL NOT NULL DEFAULT FALSE,
    ktp_photo_path VARCHAR(255) DEFAULT NULL,
    voted_at TIMESTAMP(6) WITH TIME ZONE,
    last_login TIMESTAMP(6) WITH TIME ZONE,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted BOOL NOT NULL DEFAULT FALSE
);

CREATE TABLE "pending_registrations" (
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role users_role NOT NULL,
    entity_data JSONB NOT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);

-- Set existing active users to approved verification status
UPDATE users
SET
    verification_status = 'approved'
WHERE
    is_active = TRUE;


CREATE INDEX idx_pending_registrations_user_id ON pending_registrations (user_id);
