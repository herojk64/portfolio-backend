CREATE TABLE users(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE CHECK(email <> ''),
    email TEXT NOT NULL UNIQUE CHECK(username <> ''),
    avatar TEXT NULL,
    password_hash TEXT NOT NULL,
    is_super BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
