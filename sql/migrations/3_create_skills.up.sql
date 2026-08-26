CREATE TABLE skills (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    category TEXT NULL,
    proficiency TEXT NULL,
    icon TEXT NULL,
    display_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
