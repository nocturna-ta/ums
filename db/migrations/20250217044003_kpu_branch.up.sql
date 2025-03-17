CREATE TABLE "kpu_branches"(
    id uuid NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    branch_address varchar(255) NOT NULL,
    region varchar(255) NOT NULL,
    is_active bool NOT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
)