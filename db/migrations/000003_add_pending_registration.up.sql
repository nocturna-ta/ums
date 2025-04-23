CREATE TABLE "pending_registrations" (
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role users_role NOT NULL,
    entity_data JSONB NOT NULL,
    signed_transaction TEXT,
    created_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(6) WITH TIME ZONE NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_pending_registrations_user_id ON pending_registrations (user_id);
