CREATE TABLE "users"(
    id uuid NOT NULL PRIMARY KEY,
    nik varchar(16) NOT NULL UNIQUE,
    no_telephone varchar(20),
    email varchar(100),
    name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    password_salt varchar(10) NOT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);