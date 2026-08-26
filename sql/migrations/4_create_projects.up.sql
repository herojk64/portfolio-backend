CREATE TABLE projects (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NULL,
    repo_url TEXT NULL,
    live_url TEXT NULL,
    image_url TEXT NULL,
    is_featured BOOLEAN NOT NULL DEFAULT FALSE,
    display_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE project_skills (
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    skill_id   UUID NOT NULL REFERENCES skills(id)   ON DELETE CASCADE,
    PRIMARY KEY (project_id, skill_id)
);
