CREATE TABLE experience (
    id UUID PRIMARY KEY,
    company TEXT NOT NULL,
    role TEXT NOT NULL,
    description TEXT NULL,
    start_date TEXT NOT NULL,
    end_date TEXT NULL,
    location TEXT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE experience_skills (
    experience_id UUID NOT NULL REFERENCES experience(id) ON DELETE CASCADE,
    skill_id      UUID NOT NULL REFERENCES skills(id)     ON DELETE CASCADE,
    PRIMARY KEY (experience_id, skill_id)
);
