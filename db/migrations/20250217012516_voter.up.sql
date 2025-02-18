CREATE TABLE "voters"(
    id uuid NOT NULL PRIMARY KEY,
    nik varchar(16) NOT NULL UNIQUE,
    voter_address varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    password_salt varchar(10) NOT NULL,
    is_registered bool NOT NULL DEFAULT FALSE,
    has_voted bool NOT NULL DEFAULT FALSE,
    voted_at TIMESTAMP(6) WITH TIME ZONE,
    region varchar(255) NOT NULL,
    last_login TIMESTAMP(6) WITH TIME ZONE,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);