CREATE TYPE users_role AS ENUM (
    'voter',
    'kpu_kota',
    'kpu_provinsi'
);

CREATE TABLE "users"(
    id uuid NOT NULL PRIMARY KEY,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    password_salt varchar(10) NOT NULL,
    role users_role NOT NULL,
    is_active bool NOT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);

CREATE TABLE "kpu_kota"(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    region varchar(255) NOT NULL,
    is_active bool NOT NULL,
    photo_path varchar(255) DEFAULT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);

CREATE TABLE "kpu_provinsi"(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    region varchar(255) NOT NULL,
    is_active bool NOT NULL,
    photo_path varchar(255) DEFAULT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);

CREATE TABLE "voters"(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    nik varchar(16) NOT NULL UNIQUE,
    voter_address varchar(255) NOT NULL,
    is_registered bool NOT NULL DEFAULT FALSE,
    has_voted bool NOT NULL DEFAULT FALSE,
    voted_at TIMESTAMP(6) WITH TIME ZONE,
    region varchar(255) NOT NULL,
    last_login TIMESTAMP(6) WITH TIME ZONE,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);
